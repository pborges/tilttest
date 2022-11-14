# -*- mode: Python -*-

docker_build('frontend-image', '.', dockerfile='frontend/deployments/Dockerfile')
k8s_yaml('frontend/deployments/kubernetes.yaml')
k8s_resource('frontend', port_forwards=3000)

docker_build('backend-image', '.', dockerfile='backend/Dockerfile')
k8s_yaml(helm('backend',name='backend', values='values-dev.yaml'))
k8s_resource('backend', port_forwards=8001)

docker_build('db-image', '.', dockerfile='db/deployments/Dockerfile')
k8s_yaml('db/deployments/kubernetes.yaml')
k8s_yaml('db/deployments/service.yaml')
k8s_resource('db', port_forwards=3306)

k8s_yaml('util/dns.yaml')