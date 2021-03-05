deps:
	go get -t .

startup:
	chmod 777 ./scripts/minikube-startup.sh
	./scripts/minikube-startup.sh

shutdown:
	chmod 777 ./scripts/minikube-shutdown.sh
	./scripts/minikube-shutdown.sh

build:
	docker build -t marcos30004347/k8scustomapiserver .

codegen:
	chmod 777 ./scripts/codegen.sh
	./scripts/codegen.sh

deploy:
	kubectl apply -f ./k8s/ns.yaml
	kubectl apply -f ./k8s/

undeploy:
	kubectl delete -f ./k8s/

run:
	sudo env "PATH=${PATH}" go run . --etcd-servers localhost:2379 \
    --authentication-kubeconfig ${HOME}/.kube/config \
    --authorization-kubeconfig ${HOME}/.kube/config \
    --kubeconfig ${HOME}/.kube/config