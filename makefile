deps:
	go get -t .

startup:
	chmod 777 ./hack/scripts/minikube-startup.sh
	./hack/scripts/minikube-startup.sh

shutdown:
	chmod 777 ./hack/scripts/minikube-shutdown.sh
	./hack/scripts/minikube-shutdown.sh

build:
	docker build -t marcos30004347/k8scustomapiserver .

codegen:
	chmod 777 ./hack/scripts/codegen.sh
	./hack/scripts/codegen.sh

deploy:
	kubectl apply -f ./artifacts/deploy/ns.yaml
	kubectl apply -f ./artifacts/deploy/

undeploy:
	kubectl delete -f ./artifacts/deploy/