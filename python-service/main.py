import uvicorn
from fastapi import FastAPI


from fastapi import FastAPI
from dotenv import load_dotenv, dotenv_values
from prometheus_fastapi_instrumentator import Instrumentator
import uvicorn


load_dotenv()

app = FastAPI()

Instrumentator().instrument(app).expose(app)

@app.get("/python-service/health")
def health_check():
    return {"status": "Python service is running"}


if __name__ == "__main__":
    port = dotenv_values().get("PORT", 8000)
    if port is not None:
        uvicorn.run(app, port=int(port), log_level="info", host="0.0.0.0")
