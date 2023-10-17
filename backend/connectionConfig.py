import mysql.connector

def getMysqlConnection():
    db = mysql.connector.connect(
    host = "10.8.24.11",
    user = "bitacoraTest",
    passwd = "XdEVIQShFCgK6i1UHhWb7RtcIAaZ1nNV",
    database= "bitacoraTestDB"
    )
    cursor = db.cursor()

    return cursor, db
