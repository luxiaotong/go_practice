package main

import (
	"fmt"

	"github.com/tiechui1994/gopdf"
	"github.com/tiechui1994/gopdf/core"
)

const (
	TABLE_IG = "IPAexG"
	TABLE_MD = "MPBOLD"
	TABLE_MY = "微软雅黑"
)

func main() {
	r := core.CreateReport()
	font1 := core.FontMap{
		FontName: TABLE_IG,
		FileName: "./ttf/ipaexg.ttf",
	}
	font2 := core.FontMap{
		FontName: TABLE_MD,
		FileName: "./ttf/mplus-1p-bold.ttf",
	}
	font3 := core.FontMap{
		FontName: TABLE_MY,
		FileName: "./ttf/microsoft.ttf",
	}
	r.SetFonts([]*core.FontMap{&font1, &font2, &font3})
	r.SetPage("A4", "P")

	r.RegisterExecutor(core.Executor(ComplexTableReportWithDataExecutor), core.Detail)

	r.Execute("SetDatassetApply_gopdf_Shannon.pdf")
	r.SaveAtomicCellText("SetDatassetApply_gopdf_Shannon.txt")
	fmt.Println(r.GetCurrentPageNo())
}

func ComplexTableReportWithDataExecutor(report *core.Report) {
	lineSpace := 1.0
	lineHeight := 10.0

	table := gopdf.NewTable(8, 133, 1200, lineHeight, report)
	table.SetMargin(core.Scope{})

	// 先把当前的行设置完毕, 然后才能添加单元格内容.
	c00 := table.NewCellByRange(8, 1)
	c10 := table.NewCellByRange(5, 1)
	c11 := table.NewCellByRange(3, 1)
	c20 := table.NewCellByRange(1, 1)
	c21 := table.NewCellByRange(4, 1)
	c22 := table.NewCellByRange(3, 1)
	c30 := table.NewCellByRange(1, 1)
	c31 := table.NewCellByRange(7, 1)
	c40 := table.NewCellByRange(1, 1)
	c41 := table.NewCellByRange(7, 1)

	c50 := table.NewCellByRange(1, 21)
	c51 := table.NewCellByRange(1, 7)
	c52 := table.NewCellByRange(6, 1)
	c53 := table.NewCellByRange(3, 1)
	c532 := table.NewCellByRange(3, 1)
	c54 := table.NewCellByRange(3, 1)
	c542 := table.NewCellByRange(3, 1)
	c55 := table.NewCellByRange(3, 1)
	c552 := table.NewCellByRange(3, 1)
	c56 := table.NewCellByRange(3, 1)
	c562 := table.NewCellByRange(3, 1)
	c57 := table.NewCellByRange(3, 1)
	c572 := table.NewCellByRange(3, 1)
	c58 := table.NewCellByRange(6, 1)

	d51 := table.NewCellByRange(1, 7)
	d52 := table.NewCellByRange(6, 1)
	d53 := table.NewCellByRange(3, 1)
	d532 := table.NewCellByRange(3, 1)
	d54 := table.NewCellByRange(3, 1)
	d542 := table.NewCellByRange(3, 1)
	d55 := table.NewCellByRange(3, 1)
	d552 := table.NewCellByRange(3, 1)
	d56 := table.NewCellByRange(3, 1)
	d562 := table.NewCellByRange(3, 1)
	d57 := table.NewCellByRange(3, 1)
	d572 := table.NewCellByRange(3, 1)
	d58 := table.NewCellByRange(6, 1)

	e51 := table.NewCellByRange(1, 7)
	e52 := table.NewCellByRange(6, 1)
	e53 := table.NewCellByRange(3, 1)
	e532 := table.NewCellByRange(3, 1)
	e54 := table.NewCellByRange(3, 1)
	e542 := table.NewCellByRange(3, 1)
	e55 := table.NewCellByRange(3, 1)
	e552 := table.NewCellByRange(3, 1)
	e56 := table.NewCellByRange(3, 1)
	e562 := table.NewCellByRange(3, 1)
	e57 := table.NewCellByRange(3, 1)
	e572 := table.NewCellByRange(3, 1)
	e58 := table.NewCellByRange(6, 1)

	c60 := table.NewCellByRange(1, 4)
	c61 := table.NewCellByRange(7, 1)
	c62 := table.NewCellByRange(4, 1)
	c622 := table.NewCellByRange(3, 1)
	c63 := table.NewCellByRange(4, 1)
	c632 := table.NewCellByRange(3, 1)
	c64 := table.NewCellByRange(7, 1)

	c70 := table.NewCellByRange(4, 1)
	c71 := table.NewCellByRange(4, 1)
	c80 := table.NewCellByRange(4, 1)
	c81 := table.NewCellByRange(4, 1)
	c90 := table.NewCellByRange(1, 101)
	c91 := table.NewCellByRange(1, 1)
	c92 := table.NewCellByRange(1, 1)
	c93 := table.NewCellByRange(5, 1)

	f1 := core.Font{Family: TABLE_MY, Size: 8, Style: ""}
	border := core.NewScope(4.0, 4.0, 4.0, 3.0)
	c00.SetElement(gopdf.NewTextCell(table.GetColWidth(0, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).HorizontalCentered().SetContent("数据资产确权申请书"))
	c10.SetElement(gopdf.NewTextCell(table.GetColWidth(1, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("请按照“注意事项”正确填写本表各栏"))
	c11.SetElement(gopdf.NewTextCell(table.GetColWidth(1, 5), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("此框内容由数据资产确权中心填写"))
	c20.SetElement(gopdf.NewTextCell(table.GetColWidth(2, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("③　数据资产确权名称"))
	c21.SetElement(gopdf.NewTextCell(table.GetColWidth(2, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("河南省新乡市统计局****数据"))
	c22.SetElement(gopdf.NewTextCell(table.GetColWidth(2, 5), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("①申请号0086.0451/123456789012345678.northwind.0123456789AZ\n②提交日2020-05-06T00:09:10.00+08:00"))
	c30.SetElement(gopdf.NewTextCell(table.GetColWidth(3, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("⑤　数据资产确权描述"))
	c31.SetElement(gopdf.NewTextCell(table.GetColWidth(3, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("数据资产确权描述"))
	c40.SetElement(gopdf.NewTextCell(table.GetColWidth(4, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("⑥　数据资产获取方式"))
	c41.SetElement(gopdf.NewTextCell(table.GetColWidth(4, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("□提供应用服务取得\n□机器采集取得\n☑购买取得\n□交换取得\n□其他，请注明"))

	c50.SetElement(gopdf.NewTextCell(table.GetColWidth(5, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("⑦　申请人"))
	c51.SetElement(gopdf.NewTextCell(table.GetColWidth(5, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("申请人(1)"))
	c52.SetElement(gopdf.NewTextCell(table.GetColWidth(5, 2), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("☑请求费减且已完成费减资格备案"))
	c53.SetElement(gopdf.NewTextCell(table.GetColWidth(6, 2), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("单位名称 河南省新乡市"))
	c532.SetElement(gopdf.NewTextCell(table.GetColWidth(6, 5), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("申请人类型 自然人"))
	c54.SetElement(gopdf.NewTextCell(table.GetColWidth(7, 2), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("居民身份证件号码/统一社会信用代码/组织机构代码 123456789012345678"))
	c542.SetElement(gopdf.NewTextCell(table.GetColWidth(7, 5), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("电子邮箱 aa@aa.aa"))
	c55.SetElement(gopdf.NewTextCell(table.GetColWidth(8, 2), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("国籍或注册国家（地区） 中国"))
	c552.SetElement(gopdf.NewTextCell(table.GetColWidth(8, 5), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("经常居所地或营业所所在地 北京"))
	c56.SetElement(gopdf.NewTextCell(table.GetColWidth(9, 2), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("邮政编码 123456"))
	c562.SetElement(gopdf.NewTextCell(table.GetColWidth(9, 5), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("电话 12345678901"))
	c57.SetElement(gopdf.NewTextCell(table.GetColWidth(10, 2), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("省、自治区、直辖市 河南省"))
	c572.SetElement(gopdf.NewTextCell(table.GetColWidth(10, 5), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("市县 新乡市"))
	c58.SetElement(gopdf.NewTextCell(table.GetColWidth(11, 2), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("城区（乡）、街道、门牌号 市辖区新飞大道1789号火炬园研发楼"))

	d51.SetElement(gopdf.NewTextCell(table.GetColWidth(12, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("申请人(2)"))
	d52.SetElement(gopdf.NewTextCell(table.GetColWidth(12, 2), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("□请求费减且已完成费减资格备案"))
	d53.SetElement(gopdf.NewTextCell(table.GetColWidth(13, 2), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("单位名称"))
	d54.SetElement(gopdf.NewTextCell(table.GetColWidth(14, 2), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("居民身份证件号码/统一社会信用代码/组织机构代码"))
	d55.SetElement(gopdf.NewTextCell(table.GetColWidth(15, 2), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("国籍或注册国家（地区）"))
	d56.SetElement(gopdf.NewTextCell(table.GetColWidth(16, 2), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("邮政编码"))
	d57.SetElement(gopdf.NewTextCell(table.GetColWidth(17, 2), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("省、自治区、直辖市"))
	d532.SetElement(gopdf.NewTextCell(table.GetColWidth(13, 5), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("申请人类型"))
	d542.SetElement(gopdf.NewTextCell(table.GetColWidth(14, 5), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("电子邮箱"))
	d552.SetElement(gopdf.NewTextCell(table.GetColWidth(15, 5), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("经常居所地或营业所所在地"))
	d562.SetElement(gopdf.NewTextCell(table.GetColWidth(16, 5), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("电话"))
	d572.SetElement(gopdf.NewTextCell(table.GetColWidth(17, 5), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("市县"))
	d58.SetElement(gopdf.NewTextCell(table.GetColWidth(18, 2), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("城区（乡）、街道、门牌号"))

	e51.SetElement(gopdf.NewTextCell(table.GetColWidth(19, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("申请人(3)"))
	e52.SetElement(gopdf.NewTextCell(table.GetColWidth(19, 2), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("□请求费减且已完成费减资格备案"))
	e53.SetElement(gopdf.NewTextCell(table.GetColWidth(20, 2), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("单位名称"))
	e54.SetElement(gopdf.NewTextCell(table.GetColWidth(21, 2), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("居民身份证件号码/统一社会信用代码/组织机构代码"))
	e55.SetElement(gopdf.NewTextCell(table.GetColWidth(22, 2), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("国籍或注册国家（地区）"))
	e56.SetElement(gopdf.NewTextCell(table.GetColWidth(23, 2), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("邮政编码"))
	e57.SetElement(gopdf.NewTextCell(table.GetColWidth(24, 2), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("省、自治区、直辖市"))
	e532.SetElement(gopdf.NewTextCell(table.GetColWidth(20, 5), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("申请人类型"))
	e542.SetElement(gopdf.NewTextCell(table.GetColWidth(21, 5), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("电子邮箱"))
	e552.SetElement(gopdf.NewTextCell(table.GetColWidth(22, 5), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("经常居所地或营业所所在地"))
	e562.SetElement(gopdf.NewTextCell(table.GetColWidth(23, 5), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("电话"))
	e572.SetElement(gopdf.NewTextCell(table.GetColWidth(24, 5), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("市县"))
	e58.SetElement(gopdf.NewTextCell(table.GetColWidth(25, 2), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("城区（乡）、街道、门牌号"))

	c60.SetElement(gopdf.NewTextCell(table.GetColWidth(26, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("⑧　联系人"))
	c61.SetElement(gopdf.NewTextCell(table.GetColWidth(26, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("姓名 联系人"))
	c62.SetElement(gopdf.NewTextCell(table.GetColWidth(27, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("邮政编码 123456"))
	c622.SetElement(gopdf.NewTextCell(table.GetColWidth(27, 5), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("市县 新乡市"))
	c63.SetElement(gopdf.NewTextCell(table.GetColWidth(28, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("省、自治区、直辖市 河南省"))
	c632.SetElement(gopdf.NewTextCell(table.GetColWidth(28, 5), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("居民身份证件号码/统一社会信用代码/组织机构代码"))
	c64.SetElement(gopdf.NewTextCell(table.GetColWidth(29, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("城区（乡）、街道、门牌号 市辖区新飞大道1789号火炬园研发楼"))

	c70.SetElement(gopdf.NewTextCell(table.GetColWidth(30, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("⑨ 申请文件清单\n\n申请文件清单1,申请文件清单2"))
	c71.SetElement(gopdf.NewTextCell(table.GetColWidth(30, 4), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("⑩ 附加文件清单\n\n附加文件清单1,附加文件清单2"))
	c80.SetElement(gopdf.NewTextCell(table.GetColWidth(31, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("申请人签字或者盖章\n\n申请人1"))
	c81.SetElement(gopdf.NewTextCell(table.GetColWidth(31, 4), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("数据资产确权与可信流通根中心意见\n\n查核无误，允许提交申请"))
	c90.SetElement(gopdf.NewTextCell(table.GetColWidth(32, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("④　数据资产确权编码"))
	c91.SetElement(gopdf.NewTextCell(table.GetColWidth(32, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("数据表"))
	c92.SetElement(gopdf.NewTextCell(table.GetColWidth(32, 2), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("字段"))
	c93.SetElement(gopdf.NewTextCell(table.GetColWidth(32, 3), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("哈希值"))

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			rownum := 32 + i*j
			cells := make([]*gopdf.TableCell, 3)
			cells[0] = table.NewCellByRange(1, 1)
			cells[1] = table.NewCellByRange(1, 1)
			cells[2] = table.NewCellByRange(5, 1)
			cells[0].SetElement(gopdf.NewTextCell(table.GetColWidth(rownum, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("table%d", i+1)))
			cells[1].SetElement(gopdf.NewTextCell(table.GetColWidth(rownum, 2), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("field%d", j+1)))
			cells[2].SetElement(gopdf.NewTextCell(table.GetColWidth(rownum, 3), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("12345678901234567890123456789012 12345678901234567890123456789012 12345678901234567890123456789012"))
		}
	}
	table.GenerateAtomicCell()
}
