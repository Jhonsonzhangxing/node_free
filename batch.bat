cd C:\Users\zs\go\src\github.com\sanzhang007\node_free
curl -o v2ray.txt -L "127.0.0.1/nodes?avg_i=1&avg_j=20&limit=1000000"
curl -o clash.yaml -L "127.0.0.1/clash?avg_i=1&avg_j=20&limit=1000000"
curl -o clash5-10M.yaml -L "127.0.0.1/clash?avg_i=5&avg_j=20&limit=1000000"
curl -o clash1-5M.yaml -L "127.0.0.1/clash?avg_i=1&avg_j=5&limit=1000000"
nodeFree.exe
git add ./v2ray.txt
git add ./clash.yaml
git add clash1-5M.yaml
git add clash5-10M.yaml
git add ./README.md
git add ./png/*.png
git commit -m "%date:~3,20% %time:~0,8%"
git push origin main
