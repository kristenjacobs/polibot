build:
	CGO_ENABLED=0 GOOS=linux go build -a ./...
	docker build -t kristenfjacobs/polibot:1.0.0 .

push:
	docker push kristenfjacobs/polibot:1.0.0

run:
	docker run -d -p 8081:8081 --name polibot -t kristenfjacobs/polibot:1.0.0

test:
	while true; do curl localhost:8081; sleep 1; done

clean:
	docker rm -f polibot
	rm -rf polibot

deploy:
	kubectl --kubeconfig=/Users/krisjaco/kube-admin.conf run polibot --image=kristenfjacobs/polibot:1.0.0 --replicas=2 --port=8081
	kubectl --kubeconfig=/Users/krisjaco/kube-admin.conf expose deployment polibot --name=polibot --type=NodePort

undeploy:
	kubectl --kubeconfig=/Users/krisjaco/kube-admin.conf delete deployment polibot
	kubectl --kubeconfig=/Users/krisjaco/kube-admin.conf delete services polibot

status:
	kubectl --kubeconfig=/Users/krisjaco/kube-admin.conf get deployments
	kubectl --kubeconfig=/Users/krisjaco/kube-admin.conf get pods -o wide
	kubectl --kubeconfig=/Users/krisjaco/kube-admin.conf get replicasets
	kubectl --kubeconfig=/Users/krisjaco/kube-admin.conf get services
	kubectl --kubeconfig=/Users/krisjaco/kube-admin.conf describe services polibot
