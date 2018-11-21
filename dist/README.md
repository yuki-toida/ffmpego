# Getting Started

## インストールが必要なソフトウェア[Mac]

* [Homebrew](https://brew.sh/index_ja)

```sh
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

* ffmpeg

``` sh
brew install ffmpeg
```

## ビルド方法

```sh
sh build.sh
```

* distフォルダが作成されるので

## 実行方法

* inputフォルダに連結したいmp4ファイルをコピーする

* ターミナルを開き、ffmpegoファイルがあるディレクトリに移動する

* ./ffmpego で実行する

* 正常に終了すると /output に連結されたmp4が作成される
