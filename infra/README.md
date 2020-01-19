# infra

## GKEクラスタ作成

<pre>
gcloud container clusters create clst-ftst-01 --preemptible --machine-type=f1-micro --num-nodes=3 --disk-size=10 --zone=asia-northeast1-c
</pre>

## config map

<pre>
kubectl apply -f configmap.yaml
</pre>

## deployment

### backend

<pre>
kubectl apply -f deployment_backend.yaml
</pre>

### frontend

<pre>
kubectl apply -f deployment_frontend.yaml
</pre>
