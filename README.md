# fyne2.0 中文 示例
环境: win10, go-1.18, fyne2.1.4  
 <image src="./READMES/0.png" width=154>  

## 手把手教你
* 复制字体文件到你的项目中  
  打开控制面板或在文件夹中打开 `C:\Windows\Fonts` 目录, 找到你要的字体比如 `微软雅黑`, 然后 `ctrl+c, ctrl+v` 将它复制到你的项目中  
  <image src="./READMES/1.png" width=640>  
  <image src="./READMES/2.png" width=640>  
* 使用 `fyne CLI` 打包这些字体文件  
  首先需要安装 `fyne CLI`  
  ```sh
  go install fyne.io/fyne/v2/cmd/fyne@latest
  # 输出: fyne cli version: v2.1.4
  fyne version
  ```
  这里只打包常规和粗体, 参考[这里](https://developer.fyne.io/extend/bundle)  
  ```sh
  fyne bundle msyh.ttc >> bundled.go
  # 打包粗体
  fyne bundle -append msyhbd.ttc >> bundled.go
  ```
  打包后的文件 `bundled.go` 体积有88M, 真是猛😓, 其实就是字体文件的二进制数据  
  <image src="./READMES/3.png" width=640>  
  <image src="./READMES/4.png" width=640>  
* 让你的程序使用雅黑字体, 参考[这里](https://developer.fyne.io/extend/custom-theme)  
  