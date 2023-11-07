from datetime import datetime
import pandas as pd
import numpy as np
import json

import subprocess

from pathlib import Path

import sqlite3


logs_dir = Path("")
maps_dir = Path("C:/Program Files (x86)/Steam/steamapps/common/Sid Meier's Civilization V/Maps")

working_dir = Path("maps")

mapImageExecutable = Path('Civ5MapImage.exe')
ffmpegExecutable = Path('ffmpeg.exe')


def convert_civ5_map_to_json(input_maps_dir, working_dir):
    infile = input_maps_dir / f'Civ5MapmapState_Turn000.Civ5Map'
    outfile = working_dir / f'Civ5MapmapState_Turn000.json'
    subprocess.run([
        str(mapImageExecutable),
        f'-input={infile}',
        f'-output={outfile}'
    ])
    return outfile


def merge_map_files(map_json, map2_json, out, player_colors_df, civ_df, civ_color_map_df, color_map_df):
    map_json['MapData']['CivColorOverrides'] = map2_json['MapData']['CivColorOverrides']
    for i in range(len(map_json['MapData']['CivColorOverrides'])):
        c_id = int(map_json['MapData']['CivColorOverrides'][i]['CivKey'])

        map_json['MapData']['CivColorOverrides'][i]['CivKey'] = \
        player_colors_df[player_colors_df['ID'] == c_id]['Type'].iloc[0]
        map_json['MapData']['CivColorOverrides'][i]['OuterColor']['ColorConstant'] = \
            player_colors_df[player_colors_df['ID'] == c_id]['SecondaryColor'].iloc[0]
        map_json['MapData']['CivColorOverrides'][i]['InnerColor']['ColorConstant'] = \
            player_colors_df[player_colors_df['ID'] == c_id]['PrimaryColor'].iloc[0]

    original_shape = np.array(map_json['MapData']['MapTileImprovements']).shape
    map_json['MapData']['MapTileImprovements'] = [item for sublist in map_json['MapData']['MapTileImprovements'] for
                                                  item in sublist]
    for i in range(len(map_json['MapData']['MapTileImprovements'])):
        if map2_json['MapData']['MapTileImprovements'][i]['CityId'] > 0:
            map_json['MapData']['MapTileImprovements'][i]['CityId'] = map2_json['MapData']['MapTileImprovements'][i][
                'CityId']
            map_json['MapData']['MapTileImprovements'][i]['X'] = map2_json['MapData']['MapTileImprovements'][i]['X']
            map_json['MapData']['MapTileImprovements'][i]['Y'] = map2_json['MapData']['MapTileImprovements'][i]['Y']
            map_json['MapData']['MapTileImprovements'][i]['CityName'] = map2_json['MapData']['MapTileImprovements'][i][
                'CityName']
        map_json['MapData']['MapTileImprovements'][i]['Owner'] = map2_json['MapData']['MapTileImprovements'][i]['Owner']
        map_json['MapData']['MapTileImprovements'][i]['RouteType'] = map2_json['MapData']['MapTileImprovements'][i][
            'RouteType']
    map_json['MapData']['MapTileImprovements'] = np.array(map_json['MapData']['MapTileImprovements']).reshape(
        original_shape).tolist()

    map_json['MapData']['Civ5PlayerData'] = map2_json['MapData']['Civ5PlayerData']
    map_json['MapData']['CityOwnerIndexMap'] = map2_json['MapData']['CityOwnerIndexMap']

    for i in range(len(map_json['MapData']['Civ5PlayerData'])):
        c_id = map_json['MapData']['Civ5PlayerData'][i]['CivType']
        map_json['MapData']['Civ5PlayerData'][i]['TeamColor'] = \
        civ_df[civ_df['Type'] == c_id]['DefaultPlayerColor'].iloc[0]

    for i in range(256):
        map_json['MapData']['CityOwnerIndexMap'][i] = i

    with open(out, 'w') as f:
        json.dump(map_json, f, indent=4)


def render_turn_frame(turn, injson, outpng, game_id, win_civ, win_type):
    subprocess.run([
        str(mapImageExecutable),
        '-mode=political',
        f'-turn={turn}',
        f'-input={injson}',
        f'-output={outpng}',
        f'-gameId={game_id}',
        f'-winCiv={win_civ}',
        f'-winType={win_type}'
    ])


def stitch_turn_maps_into_video(out_dir, fps, game_id):
    subprocess.run([
        str(ffmpegExecutable), '-y',
        '-framerate', str(fps),
        '-i', f'{out_dir / "map_Turn%03d.png"}',
        '-c:v', 'libx264',
        '-r', '24',
        '-crf', '30',
        '-vf', 'tpad=stop_mode=clone:stop_duration=3',
        str(out_dir / f'{game_id}.mp4')
    ])


def create_replay_frames(out_dir, total_turns, game_id, victory_civ, victory_type):
    # create dir for temporary files
    temp_dir = working_dir / game_id
    temp_dir.mkdir(parents=True, exist_ok=True)

    # Load Database things
    cnx = sqlite3.connect('Civ5DebugDatabase.db')
    player_colors_df = pd.read_sql_query("SELECT * FROM PlayerColors", cnx)
    civ_df = pd.read_sql_query("SELECT * FROM Civilizations", cnx)
    civ_color_map_df = pd.read_sql_query("SELECT * FROM PlayerColors", cnx)
    color_map_df = pd.read_sql_query("SELECT * FROM Colors", cnx)

    # Load base map file
    base_map_file = convert_civ5_map_to_json(out_dir, working_dir)
    with open(base_map_file) as f:
        base_map_json = json.load(f)

    # Create merged json for each turn
    for turn in range(total_turns):
        print(f'Processing turn {turn}/{total_turns}')

        supplementary_map_json = out_dir / f'mapStateLog_Turn{turn:03}.json'
        with open(supplementary_map_json) as f:
            supplementary_map_json = json.load(f)
        merged_json_file = working_dir / f'mapStateMerged_Turn{turn:03}.json'
        merge_map_files(base_map_json, supplementary_map_json, merged_json_file, player_colors_df, civ_df,
                        civ_color_map_df, color_map_df)

        # Use merged json for each turn to create png frames
        render_turn_frame(turn, merged_json_file, out_dir / f'map_Turn{turn:03}.png', game_id, victory_civ,
                          victory_type)


def get_result_dfs(game_result_log_path):
    result_df = pd.read_csv(game_result_log_path, skipinitialspace=True)
    cols = result_df.columns
    victory_df = result_df.copy()[cols[0:3]]

    final_score_df = result_df.copy()[cols[3:]].T.set_axis(['score'], axis=1, inplace=False).sort_values('score',
                                                                                                         ascending=False)
    final_score_df['rank'] = range(1, final_score_df.shape[0] + 1)
    final_score_df

    victory_df['game_id'] = game_result_log_path.parts[-2]
    final_score_df['game_id'] = game_result_log_path.parts[-2]

    return victory_df, final_score_df


for complete_logs_dir in Path(logs_dir).iterdir():
    print(complete_logs_dir)
    game_id = complete_logs_dir.parts[-1]
    victory_df, _ = get_result_dfs(complete_logs_dir / 'GameResult_Log.csv')
    total_turns = victory_df['Turn'][0]
    victory_civ = victory_df['VictoryCiv'][0]
    victory_type = victory_df['VictoryType'][0]

    create_replay_frames(complete_logs_dir, total_turns, game_id, victory_civ, victory_type)
    stitch_turn_maps_into_video(complete_logs_dir, 8, game_id)