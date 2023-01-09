import flask

app = flask.Flask(__name__)

@app.route("/")
def index():
    # TODO: Render the main page of the web GUI here
    return ""

@app.route("/login", methods=["GET", "POST"])
def login():
    # TODO: Handle user login here
    return ""

@app.route("/logout")
def logout():
    # TODO: Handle user logout here
    return ""

@app.route("/data")
def data():
    # TODO: Render data visualization here
    return ""

def start():
    app.run()
