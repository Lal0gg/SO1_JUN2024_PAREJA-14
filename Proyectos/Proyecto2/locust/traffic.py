import json
from random import randrange
from locust import HttpUser, between, task

debug = False

def printDebug(msg):
    if debug:
        print(msg)

class Reader():
    def __init__(self) -> None:
        self.array = []

    def pickRandom(self):
        length = len(self.array)

        if length > 0:
            random_index = randrange(0, length - 1) if length > 1 else 0
            return self.array.pop(random_index)
        else:
            print(">> Reader: No encontramos valores en el archivo")
            return None

    def load(self):
        print(">> Reader: Iniciando lectura del archivo de datos")
        try:
            with open("data.json", "r") as data_file:
                self.array = json.loads(data_file.read())
            print(f'>> Reader: Datos cargados correctamente, total de elementos: {len(self.array)}')
            print(f'>> Reader: Datos: {self.array}')
        except Exception as error:
            print(f'>> Reader: Error al leer el archivo: {error}')

class MessageTraffic(HttpUser):
    wait_time = between(0.1, 0.9)
    reader = Reader()
    reader.load()

    def on_start(self):
        print(">> MessageTraffic: Inicio de envío de tráfico")

    @task
    def PostMessage(self):
        random_data = self.reader.pickRandom()
        if random_data is not None:
            data_to_send = json.dumps(random_data)
            printDebug(data_to_send)
            resp = self.client.post("", json=random_data, headers={"Content-Type": "application/json"})
            if resp.status_code == 200:
                print(">> MessageTraffic: La petición POST se realizó correctamente.")
                print(f">> Se envió la data: {data_to_send}")
            else:
                print(f">> MessageTraffic: Error al hacer la petición POST. Código de estado: {resp.status_code}")
                print(f">> MessageTraffic: Respuesta: {resp.text}")
                self.stop(True)  # Detener la prueba si hubo un error en la petición
        else:
            print(">> MessageTraffic: Envío finalizado")
            self.stop(True)
    @task
    def GetMessage(self):
        self.client.get("/")
