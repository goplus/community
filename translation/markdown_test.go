/*
 * Copyright (c) 2023 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package translation

import (
	"fmt"
	"os"
	"testing"

	language "golang.org/x/text/language"
)

var (
	// Your api key
	mockKey = os.Getenv("NIUTRANS_API_KEY")
)

func TestTranslatePlainText(t *testing.T) {
	if mockKey == "" {
		t.Skip("NIUTRANS_API_KEY not set")
	}

	tests := []struct {
		src  string
		from string
		to   language.Tag
	}{
		{"你好", "auto", language.English},
		{"hello", "auto", language.Chinese},
	}

	trans := New(mockKey, "", "")
	for _, test := range tests {
		to, err := trans.TranslatePlainText(test.src, test.from, test.to)
		fmt.Println(to, err)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestTranslateMarkdown(t *testing.T) {
	if mockKey == "" {
		t.Skip("NIUTRANS_API_KEY not set")
	}

	tests := []struct {
		src  string
		from string
		to   language.Tag
	}{
		// {`# Hello`, "en", "zh"},
		// {`# 你好`, "zh", "en"},
		{`
# Heading

## 这是一个二级标题

这是一段普通的文本。

这是一段*粗体*的文本。

这是一段**粗体**的文本。

这是一段***粗斜体***的文本。

这是一段` + "`" + `行内代码` + "`" + `。

这是一段代码块：

` + "```" + `go
func main() {
	fmt.Println("Hello, World!")
}
` + "```" + `

这是一个列表：

- 列表项 1
- 列表项 2
	- 列表项 2.1
	- 列表项 2.2
- 列表项 3

这是一个有序列表：

1. 列表项 1
2. 列表项 2
	1. 列表项 2.1
	2. 列表项 2.2
3. 列表项 3

这是一个引用：

> 这是一段引用。

这里一个缩进代码块：

	func main() {
		fmt.Println("Hello, World!")
	}

许式伟：Go+ v1.x 的设计与实现丨Go+ 公开课 • 第一期
原创 七牛云技术团队 七牛云技术团队 2021-11-12 17:03
10 月 15 日，七牛云主办的「Go+ Together！Go+ 1.0 发布会暨 Go+ 开发者基金会启动仪式」在上海隆重召开。发布会上，七牛云 CEO、Go+ 语言发明人许式伟与 Go+ 语言贡献者共同发布了 Go+ 1.0 版本，公布了 Go+ 发展路线图。

为保持与 Go+ 开发者及行业从业者的密切沟通，共同促进、交流 Go+ 的迭代发展，Go+ 核心开发团队特策划「Go+ 公开课」系列直播分享，第一期主题为《Go+ v1.x 的设计与实现》，分享嘉宾是七牛云 CEO、Go+ 发明人 许式伟。

本文基于直播内容整理，点击文末阅读原文亦可查看完整版视频回放。

直播预告： 

11 月 13 日（周六） 19:00，Go+ 公开课 • 第 2 期正式开始！七牛云 CEO、Go+ 语言发明人许式伟将分享《import 过程与 Go+ 的模块管理》。点击按钮一键预约直播！扫码进入直播交流群，参与抽取 Go+ 官方周边~

今天的分享是 Go+ 1.0 版本后「设计与实现」系列讲座第一期，主要分享两个部分：

- Go+ 整体架构
- 如何给 Go+ 增加功能
	
第一部分主要分享 Go+ 整体架构如何实现；第二部分是和大家聊一聊如何才能够成为一个 Go+ 的贡献者，如何为 Go+ 增加新的功能。
	
一. Go+ 整体架构

Go+ 的代码比 Go 要简洁很多。我此前在知识星球「Go+ 公开课」中分享过 Go+ 的 Hello World 的三种写法：

最简洁的方式是「命令式」，比较像命令行中的代码。 

根据我的观察，小朋友在理解编程语句时，「命令式」语句的门槛是最低的，甚至比函数调用更容易被接受。因此我虽然纠结了一段时间，但还是决定将「命令式」的语法添加进去，进一步降低门槛，方便小朋友去理解。

今天我想分享的是语法背后发生的事情。下图是 Go+ 的整体架构图。我们以 Hello world 为例，来看一下执行过程会发生什么事情。

首先看左边整体的架构图。起点是 Go+ 的源代码，终点可能是两个：若运行的是可执行文件，那么终点可能是一个软件，另外则可能是一个 Package。
	
在这个过程中，其实经历了几个中间的产物，上图图形框中的部分便是对于中间产物的描述，比如 Go+ Source、Go+ Token、Go+ AST（抽象语法树）、Go DOM Writer 等。
	
其中云形框架里的「DOM Writer」模块并不是一种具象的产物，用简单的话来讲可以往里面「灌东西」，灌完便可生成 Go AST，从而转换成 Go+ 的代码。最终，通过 Go+ 的命令行程序，把它打包为一个包或可执行程序。
	
整体的流程大致就是这样。

图中的箭头上标注了很多蓝字，这些代表真正完成相关事情所需的模块。右侧是其中关键模块的名称。这些名称大家可能看起来都比较熟悉，了解编译原理的人理解起来可能会比较容易，下面我们一一来讲解一下。

1. Token & Scanner

第一步涉及的关键概念是 Token 和 Scanner。

Token 是编译原理中的一个概念，类似自然语言中「词」的概念。源代码其实就是一个字节流，它是一串文本，通过类似于词法分析的过程把字节流转换为「词」的流，这就是 Token 流。
	
这个过程在编译原理中往往叫做「词法分析」，或者叫做「lex」过程。在代码里面，我们往往把这个模块叫 Scanner，实际上就是词法分析的过程。
	
通常在我们自己写代码的时候，不需要自己来进行词法分析，一般是由 Parser 负责进行调用。所以这实际上是一个最基础的过程，把文本变成一个个单词。

比如说 if 语句的「if」是一个单词，整数「100」是一个单词，类似这样切成一个个有一定语法意义的单词，便是 Token 的概念。
	
理解了 Token 流，理解第一步就会非常容易。

2. AST & Parser

第二步涉及的关键概念是 AST & Parser。

这一步的过程是语法分析的过程，词法的上一层必然是语法，这里的逻辑和自然语言比较相近。

AST 的全称是抽象语法树，是语言的 DOM。我们在语法分析后的结果是一棵语法的树，语法树对应到自然语言中比较像一篇结构化的文章，这就是 DOM 树。DOM 模型是文本处理中很经典的模型。
	
类比一下，我们经常接触的通用文档如 XML、json 也都有自己的 DOM，这和语言有自己的 AST 实质相同。所以 AST 其实就是一个 DOM，是满足我们对语法结构化理解需要的 DOM。

Parser 其实就是语法分析，负责把 token 流或者词流转化为语法树，也就是对语法有了一个理解。
	
以上这两个过程，也就是词法分析和语法分析，在编译原理中是非常标准的过程。

下面我们看一下 Parser 的使用方式。

Go+ 中 Parser 的使用方式和 Go 是一模一样的：

fset：主要用来记录文件偏移(offset)与行列号(line:col)之间的关系，是用来定位出现语法错误时的提示，如错误出现在哪一行哪一列等相关信息；

path：最重要的一个参数，是源代码所在目录；

filter：过滤文件用途，可以传 nil。这个参数重要度不高，用来过滤掉目录中我们不希望去处理的文件；

mode：一些控制 Parser 过程的 flags，可以简单传 0；

上述是输入相关的方式，可以看到最重要的便是源代码所在的目录。输出相关的参数中，第一个参数是最重要的：
	
pkgs：得到的 AST。因为一个目录下可能有多个 pkg，所以 pkgs 是一个 map；
first：在 Parser 出错的时候，发生的第一个错误。
	
3. cl & gox

下一个阶段就是狭义的编译过程，它涉及 cl & gox 这两个模块。我们平常把全过程也叫做编译，可见这是最为关键的过程了。

编译的本质是语义分析，它负责把 Go+ AST 转换为对 gox DOM Writer 的调用。这里有一个因为兼容 Go 所以带来的复杂性——类型的推导。比如推导「变量 a」属于什么类型的变量，在语义分析中是非常难的工作。这也是 Go+ 1.0 和 Go+ 0.7 版本最大的一个差异——新版本我们引入了 gox 模块，它负责生成 Go 的抽象语法树（AST）。
	
将这个部分抽出作为一个单独的模块，是因为 Go 中的语义分析有一个很关键的点，也就是上述提到的「推导变量的类型」，这是一个非常复杂的过程，所以我们在 gox 这个 DOM Writer 组件中内置了对于变量类型的自动推导机制。

Go+ 0.7 版本实际上是在编译器中进行类型的自动推导，这使得编译器非常复杂。新版本在优化了这一机制后，给 Go+ 增加功能就会容易很多，编译器的工作也因此变得非常简单。
	
编译器输入的是 Go+ 的抽象语法树，输出的是 Go 的 DOM Writer，而 Go 生成的这个组件其实也是为了生成 Go 的抽象语法树。因此我们可以看到，cl + gox 的结果是把 Go+ AST 转化为 Go AST，本质上实现了数据格式的变换。
	
我们做个类比。了解 C++ 发明历史的朋友可能知道，C++ 最早的编译器名字叫 C-front，也就是 C 的前端。这个版本的 C++ 编译器就是把 C++ 转换成 C，再由 C 的编译器完成编译。目前 Go+ 的工作原理和 C-front 非常像，很像是一个 Go-front，也是基于这种模式我们重用了 Go 语言的工具链。
	
下面我们看下编译器的使用方式。

cl 过程一共有三个输入的参数：

pkgPath：要编译的目标，Go+ pkg 的 import 路径，也就是这个包被引用时路径名称是什么样子，这是编译过程中非常重要的一个信息；
pkg：要编译的目标 Go+ pkg 的 AST；
conf：编译用的配置。最简单就是传 nil，它会启用默认配置。
	
输出的参数一共有两个：
	
p：生成的 gox DOM Writer 的组件。这个时候返回的组件不是空的，而是已经由 cl 调用其接口完成了格式转换，已经填满了数据；
err：在编译过程中如果发生错误，则返回所有的编译错误。注意和前面 parser 过程不一样，parser 因为兼容 Go 的语义所以我们只返回第一个错误，而 cl 过程我们返回所有错误。
	
gox 可以简单分为两部分：

第一部分，用于给 cl 灌数据进而实现格式转换；

第二部分，用来生成 Go AST/Source。

第一部分的函数接口非常复杂，需要创建一个包，包里创建一个函数或者类型、变量、常量，创建函数后还需要通过 API 来编写代码。所以格式转换部分的接口数量比较多，基本通过这部分的接口可以构成一个完整 DOM 树生成的过程，这也是我们把 gox 模块叫做 DOM Writer 组件的原因，因为它大部分 API 的目的就是为了生成 DOM 的。

第二部分是在已经灌满数据后，如何生成 Go 的抽象语法树，或者说如何生成 Go 的源代码。
	
第一个函数 ASTFile 是生成 Go AST，很简单的一个函数调用。它的第一个参数是 gox DOM Writer 实例。第二个参数则比较有意思。大家知道一个目录中往往包含多个包，最常规的是一个普通代码包和一个生成测试代码的包。第二个参数便是判断我们希望生成的是哪一类语法树，然后再进行生成。

第二个函数 WriteFile 是更简单的办法，它用于生成 Go 的源代码文件，比刚才的方式就多了第一个参数，它是一个文件名。我们通过这个文件名去生成 Go 的源代码。
	
有了 Go 的源代码或者抽象语法树后，就可以调用 Go Tools 进行编译。比如我们可以执行 go run，go install，当然你也可以 go build。

当然 Go+ 发布之初我们就说过我们是双引擎的，我们既支持静态编译，也支持动态解释执行。

目前 Go+ 已经实现双引擎系统，只是两个引擎的成熟度暂时不一样。静态编译的引擎工程完成度非常高，只剩很有限的一些 Go 的功能暂时还没有支持。

但 Go+ 的解释器目前还不算太成熟，当然比我们去年发布会展示的解释器引擎要强很多。当前 Go+ 的解释器采用了开放架构，它支持非常多的  Go 的解析器作为底层引擎。比如使用 Go+ 团队自己做的 Go 解析器。

Go+ 的解析器做法目前比较简单，将 Go+ 的源代码转换成 Go+ 的抽象语法树，再变成 Go 的抽象语法树，最后通过 Go 的解析器去执行。

市面上目前包含我们自己在内的 Go 解释器完备度都不算特别高，后续我们会进行进一步的提升，尤其是在做面向数据科学领域增强的时候。

以上就是Go+整体框架的核心流程了。我们总结一下就是下面这幅图：

它是整个内容的串联。大家可以看到，整个核心过程还是非常简单的。

首先 new 一个 FlieSet，它只是错误提示的需要，真正第一步是第二行，也就是 Parser 过程。我们传一个「 . 」也就是当前目录给 Parser，Parser 完成后便得到 pkgGops，它是个包的列表。

我们假设我们关心的是 main 包，我们取出 main 包的 Go+ AST 再把它传给编译器进行编译。
	
所以第三行就是编译的过程。编译调用的是 cl.NewPackage，我们传入 Go+ 的抽象语法树，最后得到 Go 的 DOM Writer 的实例。我们调用这个实例进行 WriteFile，便可以生成 Go 的源代码。最后调用「go run」便可以执行它了。

Go+ 宏观的大框架大致如此，理解这一页的代码基本便可以把整个流程串联起来。
	
二. 如何给 Go+ 增加功能

接下来，我们从如何给 Go+ 增加功能的视角，来分析下 Go+ 的架构应该怎样理解。我们在这部分会分享通常会涉及到哪些模块以及如何进行修改。

1. 为 Go+ 新增语法

我们做一个假设：假如我们要实现 C 语言当中所谓的「三目运算符」，我们需要对哪些模块进行修改？通常我们需要进行几步判断：
	
第一步，我们需要判断是否改变了 Go+ 的抽象语法树。当然会存在不需要修改抽象语法树便可实现的功能添加，但通常来讲一个新功能往往需要对抽象语法树进行修改，比如实现三目运算符，便需要增加一个抽象语法树的节点来进行新增语言的表达。
	
修改 AST 通常就会修改两个部分：抽象语法树和语法分析。这两者是密切相关的，因为 pasrser 便是负责生成抽象语法树的组件。
	
因此新增功能时我们要做的第一件事，就是修改 ast 这个模块，为其新增一个抽象语法树的节点。仍以三目运算符为例，我们需要增加一个三目运算符表达式的节点。
	
第二步，在拥有了新的抽象语法树后，需要修改parser 模块来将三目运算符的文本转义成抽象语法树。
	
第三步，修改 cl 模块。cl 模块负责的是将新增的三目运算符表达式转换成对 gox 这个 DOM Writer 组件的调用，来生成 Go 的抽象语法树。因为三目运算符是表达式，所以我们修改的是编译器中 cl/expr.go 文件。如果新增的是语句，修改的通常为 cl/stmt.go 文件。
	
下面我们具体来看，每一个模块中的具体要修改内容。
首先，我们来看 ast 模块中我们要增加什么。
	
上图中大家可以看到「Cond？X：Y」这个文本一一对应的抽象语法树，非常容易理解。其中有三个很简单的函数是 Go+ AST 中表达式需要实现的函数，分别代表表达式的起点、终点，以及指示这是一个表达式节点的空函数。

第二部分是修改 parser，也就是语法分析模块。语法分析是负责把 token 序列转化为抽象语法树的节点，这个过程相对比较个性化，没有太多共性的关系。链接跳转的位置是三目运算符修改的代码位置，大家如果有兴趣可以试着去改一改。

第三部分是修改编译器。修改编译器的第一步不是立刻去改，而是在心中考虑清楚要增加的功能 的 AST 节点转换为 Go 的代码是什么样子的。上图中的蓝色注解便是三目运算符，它最终转换后的结果我做了一个可能的实现方式，是一个对闭包的调用。
	
这个闭包实现的难点是什么呢？是闭包的返回值类型「T」，这个 T 是什么类型光凭借语法是无法判断的，要根据语义来判断。而我们前面提到，类型推导是通过 gox 这个 DOM Writer 组件来完成，后续大家可以看到有了这个组件如何对返回值「T」进行推导。
	
如果我们用编译器也就是 cl 模块来推导返回值「T」，编译器的代码会复杂很多，而 Go+ v1.x 版本的编译器代码非常简洁，原因便是我们将类型推导的工作从编译器中解放了出去，所以它的代码基本上比较接近只进行语法分析的感觉。
	
上图中代码整体的结构非常简单，唯一的难点就是「T」如何生成。

上图是最终编辑出的编译器的部分代码。其中编译三目运算符表达式的代码是非常短的，基本只有十行左右。
	
我们来看具体是如何实现的。首先，第一行我们定义了一个自动变量，也就是需要自动推导出类型的变量。
	
第二步我们 new 了一个闭包。这个闭包的参数列表是空，所以第一个参数是 nil；返回值只有一个，所以我们创建了一个 Tuple 传入刚刚定义的变量；因为这个闭包是非不定参数的，所以第三个参数是 false；BodyStart 是开始这个闭包的函数体，开始写代码。
	
后面 DOM 的表现是非常简单的，但大家可能会觉得奇怪的一点是：为什么先编译表达式再 Return？如果学过编译原理的人会知道一个概念 —— 逆波兰表达式，也就是 DOM 所使用的表达方式，参数在前操作符在后。
	
Return 中有一个参数「1」，代表只返回一个表达式，最后的「end」表示闭包的结束。「call(0)」对应的是编译器修改部分的括号，表示闭包直接被调用并未没有传递参数.
	
差不多就是这样的概念，整体来说没有特别难理解的地方。
	
之所以说 gox 很重要，当编译器翻译 Go+ 的抽象语法树，最后变成真正能理解的代码，都需要借助 gox DOM Writer。
	
如果大家希望给 Go+ 贡献代码，理解和实现 Go+ 的功能开发，需要对于 gox 这个模块非常理解，因为基本上整体的修改过程都与 gox 相关。
	
2. 写测试案例

大家可以看到，在具备这些功能和特性后，给 Go+ 增加功能其实是非常容易的。当然我们非常重视工程化，在完成功能后一定要考虑测试案例的问题。

抽象语法树即 ast 模块部分基本是不需要测试的，因为它只是定义了数据结构，几乎没有什么代码；

parser 模块是需要测试的，需要测试是不是正确的将 Go+ 的代码转换成期望的抽象代码树。

cl 模块则需要测试是否正确将 Go+ 的代码转成期望的 Go 代码。
	
下面我们一一展开来讲解。

首先我们来看如何测试 parser 模块。对于语法分析，实际上我们基本不用写测试代码，在我们准备好的测试框架中补充测试用例即可，我觉得这是 Go+ 在生产力优化方面做得比较好好的一个点。
	
在 gop/parser 下有一个 _testdata 目录，里面都是测试用例，每一个测试用例是一个目录。我们仍以三目表达式为例，目录可以取名「ternary」。每个测试用例目录下有两个文件 —— 第一个文件 的后缀是 .gop，也就是你所撰写的代码，文件名你可以随便取，比如对三目运算符可以取 ternary.gop。另一个文件名则必须固定命名为 parser.expect，其内容是我们期望的抽象语法树 dump 成文本后的结果，用来验证我们转化后的代码是否正确。
	
对于 lambda 表达式我们准备了三个测试用例，大家可以去看一下，看完基本就可以自己使用了。
	
当然，测试用例到底是谁执行的呢？其实是 TestFromTestdata 这个函数。但这里会遇到一个问题，这个函数测试的是一批测试用例，当我们发现我们加的代码有问题想要调试时，便需要修改一下 TestFromTestdata 这个函数，将它的第一行「sel := ""」修改为「sel := "ternary"」或其他我们设置的测试用例的目录名，便可以只执行我们添加的测试用例。

接下来是如何测试 cl 模块，也就是语义分析的部分。这部分比 parser 部分更为简单。
	
大家基本上只需要在编译器的目录下找到 compile_test.go 文件，在里面增加一个测试的函数，如 TestTernaryExpr。这个函数的代码也非常简单，调用 gopClTest 函数，给出你希望 Go+ 的代码和期望生成的 Go 的代码就 OK 了。

最后，我们总结一下给 Go+ 添加功能的完整流程。
	
第一步，先 fork Go+ 的代码，然后将你 fork 的代码仓库 clone 到本地即可进行开发。

开发的第一步建议大家创建新的功能分支，然后再进行代码的修改并通过测试后，便可以提交代码，提交 pull request 给 Go+。
	
给 Go+ 添加功能的整体过程大概如此。 

通过以上添加功能的完整介绍，我想最后总结一下其中的重点内容：

增加新功能通常便是修改 gop/ast、gop/parser 和 gop/cl 三个模块
由于 gox DOM Writer 强大的类型自动推导和 Go AST 生成能力，给 Go+ 增加新功能时最复杂的 cl 过程会变得轻松很多

这两点对 Go+ 后续迭代开发非常重要。

三. 练习题 & 联系方式
	
最后，给感兴趣的朋友留几个练习题。也欢迎大家在练习的过程中，分享遇到的障碍，以便我们进一步进行过程的简化。

1. 基础练习
实现三目运算符（cond？expr1：expr2）

2. 进阶练习

特别说明：实现三目运算符的练习大家无需提交 pull request，Go+ 暂无计划添加该功能。但可以作为很好的练习题进行练习、测试。

我们也会协助大家解决训练过程中遇到的问题。联系方式如下：
	
1. Go+ 用户群（微信群）：可以直接在群里提出问题并@我，我会直接在社群进行解答；
2. Go+ 公开课（知识星球）：本次演讲的 PPT 及文字内容会同步在知识星球中，欢迎在上面提问交流。

这是一段[链接](https://www.example.com)

这是一段![图像](https://www.example.com/image.jpg)
`, "auto", language.English},
	}

	trans := New(mockKey, "", "")
	for _, test := range tests {
		translatedResult, err := trans.TranslateMarkdownText(test.src, test.from, test.to)
		fmt.Println(translatedResult, err)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestTranslateBatch(t *testing.T) {
	if mockKey == "" {
		t.Skip("NIUTRANS_API_KEY not set")
	}

	tests := []struct {
		src  []string
		from string
		to   language.Tag
	}{
		{[]string{"你好", "好的", "非常棒"}, "auto", language.English},
		{[]string{"What 's your name", "nice job"}, "auto", language.Chinese},
	}

	trans := New(mockKey, "", "")
	for _, test := range tests {
		translatedResult, err := trans.TranslateBatchPlain(test.src, test.from, test.to)
		fmt.Println(translatedResult, err)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestTranslateEmpty(t *testing.T) {
	if mockKey == "" {
		t.Skip("NIUTRANS_API_KEY not set")
	}

	tests := []struct {
		src  string
		from string
		to   language.Tag
	}{
		{"", "auto", language.English},
		{"", "auto", language.Chinese},
	}

	trans := New(mockKey, "", "")
	for _, test := range tests {
		translatedResult, err := trans.TranslateMarkdownText(test.src, test.from, test.to)
		fmt.Println(translatedResult, err)
		if err != nil {
			t.Fatal(err)
		}
	}
}
