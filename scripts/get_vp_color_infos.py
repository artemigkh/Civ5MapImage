import sqlite3


cnx = sqlite3.connect('Civ5DebugDatabase.db')
civ_color_map_df = pd.read_sql_query("SELECT * FROM PlayerColors", cnx)

color_map_df = pd.read_sql_query("SELECT * FROM Colors", cnx)

for index, row in color_map_df.iterrows():
    print(f'colorMap["{row["Type"]}"] = color.RGBA{"{"}{int(row["Red"]*255)}, {int(row["Green"]*255)}, {int(row["Blue"]*255)}, {int(row["Alpha"]*255)}{"}"}')


for index, row in civ_color_map_df.iterrows():
    print(
f'''
civColorMap["{row['Type']}"] = CivColor{'{'}
    OuterColor: colorMap["{row['PrimaryColor']}"],
    InnerColor: colorMap["{row['SecondaryColor']}"],
    TextColor:  colorMap["{row['TextColor']}"],
{'}'}''')
