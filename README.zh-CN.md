<p align="center">
<img src="./assets/theme_default_cool.png" alt="logo" width="300" height="120" style="border-radius: 12px;" />
</p>

<h3 align="center">CoolCaptcha</h3>

## 介绍

`CoolCaptcha`是一款我自认为很酷的图形验证码。可以很快速便捷地生成验证码图片的base64数据，并且，该库在很多配置来自定义样式，使其符合产品的风格。

（小声bb：这真的是泰裤辣！）



语言: [简体中文](https://github.com/Kuari/CoolCaptcha/README.zh-CN.md) English



## 案例

![theme_default_cool](./assets/theme_default_cool.png)![theme_default_random](./assets/theme_default_random.png)

![theme1_cool](./assets/theme1_cool.png)![theme1_random](./assets/theme1_random.png)

![theme2_cool](./assets/theme2_cool.png)![theme2_random](./assets/theme2_random.png)





## 安装

```go
go get github.com/Kuari/coolCaptcha
```



## 使用

### 快速生成

快速生成验证码图片，无需任何配置

```go
package main

import (
	"github.com/Kuari/coolCaptcha"
)

func main() {
	imageBase64Data, code, err := coolCaptcha.New().Generate()
}
```

### 自定义配置

```go
package main

import (
	"github.com/Kuari/coolCaptcha"
)

func main() {
	options := []coolCaptcha.Options{
		coolCaptcha.SetBackgroundHexColor("#c4e1f6"),                            // 设置图片背景色
		coolCaptcha.SetFontHexColor("#312E2E"),                                  // 设置字体颜色
		coolCaptcha.SetLineHexColors([]string{"#f596a1", "#fadeeb", "#f9c975"}), // 设置线条颜色, 会从中随机选择3条, 因此该参数至少设置3个值
		coolCaptcha.SetWidth(300),                                               // 设置图片的宽度
		coolCaptcha.SetHeight(120),                                              // 设置图片的高度
		coolCaptcha.SetCodeType(coolCaptcha.NumericCharacters),                  // 设置验证字符的类型, 有UppercaseEnglishCharacters, NumericCharacters, MixedCharacters三个类型
		coolCaptcha.SetDevMode(true),                                            // 设置开发模块, 适用于开发时将base64数据保存为图片, 便于查看生成效果
	}

	imageBase64Data, code, err := coolCaptcha.New(options...).Generate()
}
```

### 自定义验证码

可以通过自己生成的验证码来生成图片

```go
package main

import (
	"github.com/Kuari/coolCaptcha"
)

func main() {
	// customCode方法仅支持4位字符的英文和数字,
	// 当传入英文时, 将会被大写然后再使用, 因此, 当使用自定义字符时, 输出的code都是大写的, 验证的时候请务必注意
	// 全部大写化是为了小写英文字母与数字的歧义
	imageBase64Data, code, err := coolCaptcha.New().CustomCode("cool").Generate()
}
```



## Q&A

### 1. 默认字体可以免费商用吗？

默认字体使用的是 [blowbrush](https://www.dafont.com/blowbrush.font)，其是免费商用的，可以放心使用。

后续将开放自定义字体设置，敬请期待。

### 2. OCR能够破解吗？

`CoolCaptcha`存在被OCR破解的风险，请结合具体场景使用。

### 3. 如何避免英文字母与数字区分不清？

`CoolCaptcha`将英文全部大写化，因此对于英文字母与数字之间较为清晰。还可以通过`SetCodeType`来设置验证字符为纯英文字符或者纯数字。

### 4. 为什么要开发这个库？

遇到一些具体场景，需要用到图形验证码，但是发现几个问题。首先是当前的验证码大都样式相识，我觉得用户体验很重要，设计师和前端辛辛苦苦搞的产品，有一个风格不同的图形验证码会很奇怪。其次是go的图形验证码库有点少，之前跟一个java开发工程师合作，很羡慕那种快速输出一个图形验证码的能力。因此，就搞起来了。