import mysql.connector

def databaseConnect():
    db = mysql.connector.connect(
        host = "127.0.0.1",
        user = "bitacoraU",
        passwd = "test-passwd",
        database= "bitacoraDB"
    )
    return db
