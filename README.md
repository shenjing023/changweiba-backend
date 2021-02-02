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
  -  Install Kiali dashboard, along with Prometheus, Grafana, and Jaeger [kiali](https://istio.io/latest/docs/setup/getting-started/#dashboard) [grafana](https://istio.io/latest/docs/tasks/observability/metrics/using-istio-dashboard/)
  -  Install loki

## Doc

  - 仅使用k8s的service在grpc服务时不能负载均衡，需要配合istio的virtualservice和destinationrule使用 [这里](https://medium.com/getamis/istio-%E5%9F%BA%E7%A4%8E-grpc-%E8%B2%A0%E8%BC%89%E5%9D%87%E8%A1%A1-d4be0d49ee07) [github](https://github.com/alanchchen/grpc-lb-istio)
