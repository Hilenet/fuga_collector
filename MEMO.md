
## repositoryパターン
domain/repository

repositoryパターンは永続化の抽象化
-> 細部は残す必要なし
-> entity以外(中間テーブル)は直でいじらせる?

## usecase interacterがDIableなのは何故
interecterを差し替えることで，controller以下を分離可能
-> controllerテストのスタブ，
presenterへの出力についても同様
-> presenterからAPI側のモック

## inputData欲しい
受け渡す段階で変換？

## したい
depからgo modulesに移行
* 1.12から公式
* 1.13から標準有効化予定

## controllerの単位
slack.\*Eventに対応するhandleMethod生やして，usecaseを中で行う．
ex) addInfo1EventとaddInfo2EventをマージしてupdateのUsecaseでhandle

そうなるとinfraを上から書いた方が良さげ

## wip....
controller/handler書いてる途中
* 外部キーの解決をいつやる
  * Usecaseに外部のrepo持たせてFind生やし，Controller段階で解決
  * DTOでUsecaseに渡し，Usecaseで解決
  * そもそもentityに持たせるの生のIDで良いのでは
* 中間テーブルの取り扱い
  * モデルは持ってない
  * 変更はUserあたりで行うか
  * entityとしてモデル作ってしまうか


