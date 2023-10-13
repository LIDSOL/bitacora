import mysql.connector

def getMysqlConnection():
    db = mysql.connector.connect(
        host = "127.0.0.1",
        user = "bitacoraU",
        passwd = "test-passwd",
        database= "bitacoraDB"
    )
    cursor = db.cursor()

    return cursor, db
