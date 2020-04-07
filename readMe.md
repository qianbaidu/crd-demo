# crd demo

- 环境需求
    - kubebuilder
    - kustomize
    - k8s
    - kubectl
    - golang
    - docker


- 运行

```
# 构建运行
make && make install && make run
# 创建资源
kubectl apply -f config/samples/
# 查看
➜  kubectl get po -n crd-demo-system
➜  kubectl get deploy -n crd-demo-system
```