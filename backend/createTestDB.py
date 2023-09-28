import requests
import random
from faker import Faker

fake = Faker()

# Define the URL where you are making the POST requests
url = 'http://127.0.0.1:5000/'

# Users
for i in range(1, 20):
    name = fake.first_name()
    surname = fake.last_name()

    user = {
        'userID': str(i),
        'name': name,
        'surname': surname,
        'userType': 'S',
        'email': 'test@mail.com'
    }

    x = requests.post(url + 'addUser', json=user)
    print(x.text)

# Projects
for i in range(1, 10):
    project_name = fake.text(max_nb_chars=10)
    manager = random.randint(1,20)
    description = fake.text(max_nb_chars=200)

    project = {
        'name': project_name,
        'manager': str(manager),
        'description': description
    }

    x = requests.post(url + 'addProject', json=project)
    print(x.text)

# Add log
for i in range(1, 20):
    userID = random.randint(1,20)
    projectID = random.randint(1,10)

    log = {
        'userID': str(userID),
        'projectID': str(projectID)
    }

    x = requests.post(url + 'addLog', json=log)
    print(x.text)
