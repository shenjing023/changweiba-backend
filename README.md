# changweiba-backend
<p align="center">
  <img src="https://github.com/shenjing023/changweiba-backend/blob/master/logo.jpg"/>
  <br><strong><a href="https://github.com/shenjing023/changweiba-backend" target="_blank">消灭权限暴政，世界属于肠胃</a></strong>
</p>

## Install steps

  -  Install [Docker](https://www.docker.com/)   
     Install Docker if you don't have it installed
  -  Install [kind](https://kind.sigs.k8s.io/)   
     If you already have a K8S cluster, skip it. Install kind, and then create cluster, like this 
     ```bash
     kind create cluster --name changweiba
     ```
  -  Install [Istio](https://istio.io/latest/)   
     Install istio follow the official documentation, and set automatic sidecar injection 
     ```bash
     kubectl label namespace default istio-injection=enabled
     ```
     and then 
     ```bash
     istioctl install
     ```
  -  Set service config in yaml config file
  -  Run service
     Apply each service yaml file, and finally apply changweiba-gateway.yaml
     ```bash
     kubectl apply -f changweiba-gateway.yaml
     ``` 
  -  Install [Kiali dashboard](https://istio.io/latest/docs/setup/getting-started/#dashboard), along with Prometheus, [Grafana](https://istio.io/latest/docs/tasks/observability/metrics/using-istio-dashboard/), and Jaeger
     ```bash
     // move to istio directory
     cd istio-1.9.0
     kubectl apply -f samples/addons
     kubectl rollout status deployment/kiali -n istio-system
     ```
     and use port forwarding to access grafana in cluster
     ```bash
     kubectl -n istio-system port-forward $(kubectl -n istio-system get pod -l app=grafana -o jsonpath='{.items[0].metadata.name}') 3000:3000
     ```
     or
     ```bash
     istioctl dashboard grafana
     ```
  -  Install loki
     ```bash
     kubectl apply -f loki-deploy.yaml
     ```
      

## Doc

  - 仅使用k8s的service在grpc服务时不能负载均衡，需要配合istio的virtualservice和destinationrule使用 [这里](https://medium.com/getamis/istio-%E5%9F%BA%E7%A4%8E-grpc-%E8%B2%A0%E8%BC%89%E5%9D%87%E8%A1%A1-d4be0d49ee07) [github](https://github.com/alanchchen/grpc-lb-istio)
  - promtail error:
    ```bash
    Failed to list *v1.Endpoints: endpoints is forbidden: User "system:serviceaccount:istio-system:default" cannot list resource "endpoints" in API group "" in the namespace "default"
    ```
    [here](https://github.com/prometheus-operator/prometheus-operator/issues/2155#issuecomment-441002864)
    solution:
    ```bash
    kubectl create clusterrolebinding default-view --clusterrole=view --serviceaccount=istio-system:default
    ```
