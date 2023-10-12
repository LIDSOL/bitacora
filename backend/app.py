from flask import Flask, jsonify, request
from flask_cors import CORS
import mysql.connector

# Database connection 
db = mysql.connector.connect(
    host = "127.0.0.1",
    user = "bitacoraU",
    passwd = "test-passwd",
    database= "bitacoraDB"
)

def userExists(userID):
    cursor = db.cursor()

    sql = "SELECT * FROM users WHERE userid = %s;"
    val = (userID, )

    cursor.execute(sql, val)
    res = cursor.fetchall()

    if res:
        return True
    else:
        return False
 
def projectExists(projectID):
    cursor = db.cursor()

    sql = "SELECT * FROM projects WHERE id = %s;"
    val = (projectID, )

    cursor.execute(sql, val)
    res = cursor.fetchall()

    if res:
        return True
    else:
        return False

app = Flask(__name__)
CORS(app)

# Testing Route
@app.route('/ping', methods=['GET'])
def ping():
    return jsonify({'response': 'pong!'})

@app.route('/listProjects', methods=['GET'])
def listProjects():
    cursor = db.cursor()

    sql = "SELECT id,name FROM projects"
    cursor.execute(sql)

    projects =  cursor.fetchall()

    res = dict()

    for p in projects:
        key = p[0]
        val = p[1]
        res[key] = val

    return jsonify(res)

@app.route('/userExists', methods=['POST'])
def userExistsW():
    cursor = db.cursor()

    userID = request.json['userID']

    if userExists(userID):
        return jsonify({'response': 'true'})
    else:
        return jsonify({'response': 'false'})

@app.route('/projectExists', methods=['POST'])
def projectExistsW():
    cursor = db.cursor()

    projectID = request.json['projectID']

    if projectExists(projectID):
        return jsonify({'response': 'true'})
    else:
        return jsonify({'response': 'false'})

@app.route('/addUser', methods=['POST'])
def addUserW():
    cursor = db.cursor()

    userID = request.json['userID']
    name = request.json['name']
    surname = request.json['surname']
    userType = request.json['userType']
    email = request.json['email']

    if userExists(userID):
        return jsonify({'response': 'User already exists'})
    else:
        cursor = db.cursor()

        sql = "INSERT INTO users (userid, name, surname, userType, email) VALUES (%s, %s, %s, %s, %s)"
        val = (userID, name, surname, userType, email)

        cursor.execute(sql, val)
        db.commit()
        return jsonify({'response': 'ok'})

@app.route('/addProject', methods=['POST'])
def addProjectW():
    cursor = db.cursor()

    name = request.json['name']
    manager = request.json['manager']
    description = request.json['description']


    sql = "INSERT INTO projects (name, manager, description) VALUES (%s, %s, %s)"
    val = (name, manager, description)

    cursor.execute(sql, val)
    db.commit()
    return jsonify({'response': 'ok'})


@app.route('/addLog', methods=['POST'])
def addLogW():
    userID = request.json['userID']
    projectID = request.json['projectID']

    if userExists(userID) and projectExists(projectID):
        cursor = db.cursor()

        sql = "INSERT INTO logs (userID, projectID) VALUES (%s, %s)"
        val = (userID, projectID)

        cursor.execute(sql, val)
        db.commit()
        return jsonify({'response': 'ok'})
    elif not userExists(userID):
        return jsonify({'response': 'User is missing'})
    elif not projectExists(projectID):
        return jsonify({'response': 'Project is missing'})
