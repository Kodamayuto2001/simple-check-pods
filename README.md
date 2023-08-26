# ローカルで試す。

##  まずはコンテキストの確認
`kubectl config current-context`

##  対象のコンテキストを選択
`kubectl config use-context docker-desktop`

##  事前に作成
```zsh
cd sample-manifests
kubectl apply -f simple-deployment-manifest.yaml
```

##  手動で確認
```zsh
kubectl get pods
```

##  プログラムから確認
```zsh
go run hello.go
```

##  問題がないことを確認して実験が終了
- [x] 問題ない