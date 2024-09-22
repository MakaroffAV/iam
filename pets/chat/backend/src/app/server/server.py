import json

from flask import Flask
from flask import request
from flask import Response
from flask_cors import CORS
from flask import send_from_directory

from modules.email import Email as ModuleEmail
from modules.queue import Queue as ModuleQueue

from src.config.email import Email as ConfigEmail
from src.config.db import Database as ConfigDatabase
from src.config.broker import Broker as ConfigBroker
from src.config.chroma import ChromaDB as ConfigChromaDB
from src.config.embedding import Embedding as ConfigEmbedding

from src.repo.user import User as RepoUser
from src.repo.code import Code as RepoCode
from src.repo.token import Token as RepoToken
from src.repo.project import Project as RepoProject
from src.repo.project_file import ProjectFile as RepoProjectFile
from src.repo.project_query import ProjectQuery as RepoProjectQuery

from src.service.login import Login as ServiceLogin
from src.service.token import Token as ServiceToken
from src.service.project import Project as ServiceProject
from src.service.project_file import ProjectFile as ServiceProjectFile

from src.handlers.login import Login as HandlerLogin
from src.handlers.token import Token as HandlerToken
from src.handlers.project import Project as HandlerProject
from src.handlers.project_file import ProjectFile as HandlerProjectFile


app = Flask(__name__)
CORS(app)
# CORS(app, origins=[<host>])

sv_project_file = ServiceProjectFile(
    RepoProjectFile(ConfigDatabase()),
)
sv_token = ServiceToken(
    RepoCode(ConfigDatabase()), RepoToken(ConfigDatabase()),
)
sv_login = ServiceLogin(
    RepoUser(ConfigDatabase()), RepoCode(ConfigDatabase()), ModuleEmail(ConfigEmail()),
)
sv_project = ServiceProject(
    ModuleQueue(ConfigBroker()), RepoToken(ConfigDatabase()), RepoProject(ConfigDatabase()), sv_project_file, RepoProjectQuery(ConfigDatabase()))


@app.route("/api/login", methods=["POST"])
def login():
    return Response(
        **HandlerLogin(sv_login).post(
            request.form.get("email", None),
        ),
    )


@app.route("/api/token/get", methods=["POST"])
def token_get():
    return Response(
        **HandlerToken(sv_token).create(
            request.form.get("code", None),
            request.form.get("uuid", None),
        )
    )


@app.route("/api/token/check", methods=["POST"])
def token_check():
    return Response(
        **HandlerToken(sv_token).check(
            request.form.get("token", None),
        ),
    )


@app.route("/api/project/all", methods=["POST"])
def project_all():
    return Response(
        **HandlerProject(sv_project).all(
            request.form.get("token",  None),
            request.form.get("status", None),
        )
    )


@app.route("/api/project/get", methods=["POST"])
def project_get():
    return Response(
        **HandlerProject(sv_project).get(
            request.form.get("token", None),
            request.form.get("project_uuid", None),
        )
    )


@app.route("/api/project/new", methods=["POST"])
def project_new():
    return Response(
        **HandlerProject(sv_project).create(
            request.form.get("name",  None),
            request.form.get("desc",  None),
            request.form.get("token", None), request.files.getlist("file"),
        )
    )


@app.route("/api/project/query", methods=["POST"])
def project_query():
    return Response(
        **HandlerProject(sv_project).new_query(
            request.form.get("token", None),
            request.form.get("query", None),
            request.form.get("enrichment", None),
            request.form.get("project_uuid", None)
        )
    )


@app.route("/api/project/file/list", methods=["POST"])
def project_file_list():
    return Response(
        **HandlerProjectFile(sv_project_file).list(
            request.form.get("token", None),
            request.form.get("project_uuid", None)
        )
    )


@app.route("/api/project/query/history", methods=["POST"])
def project_query_history():
    return Response(
        **HandlerProject(sv_project).query_history(
            request.form.get("token", None),
            request.form.get("project_uuid", None),
        )
    )


@app.route("/api/project/query/result", methods=["GET"])
def project_query_result():
    query_uuid = request.args.get("uuid", None)
    if query_uuid:
        return send_from_directory(
            directory="/Users/makarov_aleksei/Desktop/iam/pets/chat/backend/queries", path="{}.docx".format(query_uuid), as_attachment=True,
        )
    else:
        return Response(
            response=json.dumps({"reason": "undefined project uuid"}), status=400, mimetype="application/json"
        )


def server():
    app.run(debug=True, port=5001, host="0.0.0.0")
