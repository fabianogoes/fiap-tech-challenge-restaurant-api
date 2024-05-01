import docker
import os
from datetime import datetime

# Nome do repositório no Docker Hub
repository_name = "fabianogoes/restaurant-api"

# Número máximo de tags a serem mantidas
max_tags = 3

# Autenticação no Docker Hub
client = docker.DockerClient(base_url='unix://var/run/docker.sock')  # Use a API local do Docker
client.login(username=os.environ['DOCKER_USERNAME'], password=os.environ['DOCKER_PASSWORD'])

# Obter lista de tags do repositório
tags = client.api.tags(repository_name)

# Ordenar as tags por data
tags.sort(key=lambda x: datetime.strptime(x['last_updated'], "%Y-%m-%dT%H:%M:%S.%fZ"), reverse=True)

# Excluir tags extras
for tag in tags[max_tags:]:
    client.api.delete(f"/v2/repositories/{repository_name}/tags/{tag['name']}")

print(f"As seguintes tags foram mantidas em {repository_name}:")
for tag in tags[:max_tags]:
    print(tag['name'])
