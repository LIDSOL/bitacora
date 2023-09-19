from flask import Flask, render_template
import mysql.connector

def databaseConnect():
    db = mysql.connector.connect(
        host = "127.0.0.1",
        user = "bitacoraU",
        passwd = "test-passwd",
        database= "bitacoraDB"
    )
    return db

def userExists(db, userID):
    cursor = db.cursor()

    sql = "SELECT * FROM users WHERE userid = %s;"
    val = (userID, )

    cursor.execute(sql, val)
    res = cursor.fetchall()

    if res:
        return True
    else:
        return False
 
def projectExists(db, projectID):
    cursor = db.cursor()

    sql = "SELECT * FROM projects WHERE id = %s;"
    val = (projectID, )

    cursor.execute(sql, val)
    res = cursor.fetchall()

    if res:
        return True
    else:
        return False

def addUser(db, userID, name, surname, userType, email):
    if userExists(db, userID):
        return False
    else:
        cursor = db.cursor()

        sql = "INSERT INTO users (userid, name, surname, userType, email) VALUES (%s, %s, %s, %s, %s)"
        val = (userID, name, surname, userType, email)

        cursor.execute(sql, val)
        db.commit()
        return True

def addProject(db, name, manager, description):
    cursor = db.cursor()

    sql = "INSERT INTO projects (name, manager, description) VALUES (%s, %s, %s)"
    val = (name, manager, description)

    cursor.execute(sql, val)
    db.commit()

def addLog(db, userID, projectID):
    if userExists(db, userID) and projectExists(db, projectID):
        cursor = db.cursor()

        sql = "INSERT INTO logs (userID, projectID) VALUES (%s, %s)"
        val = (userID, projectID)

        cursor.execute(sql, val)
        db.commit()
        return True
    else:
        return False

def listProjects(db):
    cursor = db.cursor()

    sql = "SELECT id,name FROM projects"
    cursor.execute(sql)

    return cursor.fetchall()

if __name__=="__main__":
    db = databaseConnect()
    #print(addUser(db, "123456", "John", "Doe", "S", "example@email.net"))
    #addProject(db, "Bitacora", "123456", "Creacion de bitacora para administracion de entradas al laboratorio")
    #addLog(db, "123456","1")
    #print(userExists(db, "123456"))
    #print(listProjects(db))
    #print(projectExists(db, 1))


app = Flask(__name__)