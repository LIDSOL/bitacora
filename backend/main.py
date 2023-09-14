# Python implementation to create a Database in MySQL
import dbconfig

# Add user to database
def addUser(db, userID, name, surname, userType, email):
    cursor = db.cursor()

    sql = "INSERT INTO users (userid, name, surname, userType, email) VALUES (%s, %s, %s, %s, %s)"
    val = (userID, name, surname, userType, email)

    cursor.execute(sql, val)
    db.commit()

def addProject(db, name, manager, description):
    cursor = db.cursor()

    sql = "INSERT INTO projects (name, manager, description) VALUES (%s, %s, %s)"
    val = (name, manager, description)

    cursor.execute(sql, val)
    db.commit()

def addLog(db, userID, projectID):
    cursor = db.cursor()

    sql = "INSERT INTO logs (userID, projectID) VALUES (%s, %s)"
    val = (userID, projectID)

    cursor.execute(sql, val)
    db.commit()

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

if __name__=="__main__":
    db = dbconfig.databaseConnect()
    #addUser(db, "123456", "John", "Doe", "S", "example@email.net")
    #addProject(db, "Bitacora", "123456", "Creacion de bitacora para administracion de entradas al laboratorio")
    #addLog(db, "123456","1")
    #print(userExists(db, "123456"))
