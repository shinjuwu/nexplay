"""Mock Game Server for local development.
Responds to platform backend's notification API calls.
"""
from http.server import HTTPServer, BaseHTTPRequestHandler
import json

# Default killdive data matching expected game IDs
KILLDIVE_DATA = json.dumps([
    {"AgentId": 0, "GameId": 1001, "RoomId": 10010, "Killrate": 0.05, "Newkillrate": 0.03, "Activenum": 100},
    {"AgentId": 0, "GameId": 1002, "RoomId": 10020, "Killrate": 0.05, "Newkillrate": 0.03, "Activenum": 100},
    {"AgentId": 0, "GameId": 1003, "RoomId": 10030, "Killrate": 0.05, "Newkillrate": 0.03, "Activenum": 100},
    {"AgentId": 0, "GameId": 1004, "RoomId": 10040, "Killrate": 0.05, "Newkillrate": 0.03, "Activenum": 100},
    {"AgentId": 0, "GameId": 1005, "RoomId": 10050, "Killrate": 0.05, "Newkillrate": 0.03, "Activenum": 100},
    {"AgentId": 0, "GameId": 2001, "RoomId": 20010, "Killrate": 0.05, "Newkillrate": 0.03, "Activenum": 100},
    {"AgentId": 0, "GameId": 2002, "RoomId": 20020, "Killrate": 0.05, "Newkillrate": 0.03, "Activenum": 100},
    {"AgentId": 0, "GameId": 3001, "RoomId": 30010, "Killrate": 0.05, "Newkillrate": 0.03, "Activenum": 100},
])

GAMESETTING_DATA = json.dumps([])


class MockHandler(BaseHTTPRequestHandler):
    def _respond(self, data):
        self.send_response(200)
        self.send_header("Content-Type", "application/json")
        self.end_headers()
        self.wfile.write(json.dumps(data).encode())

    def do_GET(self):
        path = self.path.split("?")[0].strip("/")
        print(f"GET /{path}")

        if path == "getdefaultkilldiveinfo":
            self._respond({"code": 0, "data": KILLDIVE_DATA})
        elif path == "getdefaultgamesetting":
            self._respond({"code": 0, "data": GAMESETTING_DATA})
        elif path == "querygold":
            self._respond({"code": 0, "gold": 0, "lockgold": 0})
        elif path == "querypluralgold":
            self._respond({"code": 0, "data": []})
        elif path == "getgameusernewbie":
            self._respond({"code": 0, "data": json.dumps({"data": [], "total_newbie_limit": 0})})
        elif path == "getrealtimegameratio":
            self._respond({"code": 0, "data": "[]"})
        elif path == "getplinkoballmaxodds":
            self._respond({"code": 0, "data": "[]"})
        else:
            self._respond({"code": 0})

    def do_POST(self):
        content_length = int(self.headers.get("Content-Length", 0))
        body = self.rfile.read(content_length) if content_length else b""
        path = self.path.split("?")[0].strip("/")
        print(f"POST /{path} body={body.decode('utf-8', errors='replace')[:200]}")
        self._respond({"code": 0})


if __name__ == "__main__":
    port = 9642
    print(f"Mock Game Server listening on port {port}")
    HTTPServer(("0.0.0.0", port), MockHandler).serve_forever()
