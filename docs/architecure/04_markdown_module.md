
# Markdown module

## Module purpose

 markdown模块使用cherry-markdown实现了一个markdown在线编辑器，左侧是编辑区，右侧是预览区。添加了对go+语言的高亮渲染。

## Module scope

模块输出用户编辑的markdown文本，输入为已经存在用户markdown文本。

## Module structure

在这里详细描述模块的内部结构和行为。

## Module Interface

在这里提供模块的公开接口的详细信息，包括函数名称、参数、返回值和可能的错误。
submit_markdown （）{

    return  {
        FormData {
            title : String,
            content: String,
            trans: String,
        }
    }
    
}

## Functions

### gop渲染

示例：

 ```gop
  
   // go+代码
   println ""
 ```

### submit_markdown

- 功能：提交用户编辑的markdown文档
- 入参：无
- 返回：无
- 错误：axios请求异常




