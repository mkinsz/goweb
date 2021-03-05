模块（Module）、组件（Component）、包（Package），这些概念对于我们技术同学并不陌生，但并不是所有人都能理解其要义。

深入理解之后，我才发现，其背后的深意是分类思维。而这种分类也是应用架构的核心所在，通过不同粒度、不同层次的分类，把复杂的软件系统实现控制在可以被理解、被维护的程度。否则，对于动则上100万行代码的软件，人类根本没有办法理解和维护。

应用（Application）：应用系统，有多个Module组成，用方框表示。

模块（Module）：一个Module是有一组Component构成，用正方体表示。

组件（Component）：表示一个可以独立提供某方面功能的物件，用UML的组件图表示。

包（Package)：Package相对比较tricky，它是一种组织形式，和粒度不是一个维度的，也就是说，一个Component可以包含多个Package，一个Package也可以包含多个Component。


## 应用架构的要素

Architecture = Structure Of Components + Relationships + Principles & Guidelines

即架构是一种结构，是由物件（Components）+ 物件之间的关系 + 指导原则组成的。

应用架构也是如此，从大的层面来说，企业级应用都逃不过如下图所示的三层结构，即前端、后端和数据库。

