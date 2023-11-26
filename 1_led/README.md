# 1.LEDの点滅

LEDを点滅させるソースコードです。
いわゆるLチカです。

# GPIO

<img width="300" alt="image" src="https://github.com/hinagithub/raspberrypi-sensor/assets/44778704/af956fa3-ca6a-4a68-be2b-5e3f88b304a0">

ランプには上から順に
- R = Red (黄色のコード)
- G = Green(赤のコード)
- GND = Ground(黒のコード)
という接続ピンがあります。


Greenなのに赤いコードなのが紛らわしいので注意です。

Groundは、電気回路において基準点として使用される参照電位を指します。

グラウンドは通常、電流が戻る経路として機能します。

電気回路が正しく動作するためには不可欠な要素らしいです。

# 実行方法

1_LEDディレクトリ配下へ移動

実行

```
go run main.go
```


停止
```
Ctrl+c
```
