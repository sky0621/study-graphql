# infra

## GKEクラスタ作成

<pre>
gcloud container clusters create clst-ftst-03 --preemptible --machine-type=f1-micro --num-nodes=3 --disk-size=10 --zone=asia-northeast1-c
</pre>

## auth

<pre>
gcloud container clusters get-credentials clst-ftst-03 --zone asia-northeast1-c --project [GCP Project]
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

## config

### current

<pre>
kubectl config current-context
</pre>

### list

<pre>
kubectl config get-contexts
</pre>

### set

<pre>
kubectl config use-context [context name]
</pre>

### delete

<pre>
kubectl config delete-context [context name]
</pre>
