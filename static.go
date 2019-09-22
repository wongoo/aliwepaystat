// Copyright 2019 wongoo. All rights reserved.

// Code generated by "makestatic"; DO NOT EDIT.

package aliwepaystat

var Files = map[string]string{
	"layout.html": `{{ define "layout" }}
    <!DOCTYPE html>
    <html lang="zh">
    <head>
        <meta charset="UTF-8">
        <title>statistics report</title>
        <style type="text/css">
            body{font-size: 12px;}
            table{border: 1px solid #ccc; border-collapse: collapse;}
            th,td{border:1px solid #bbb;padding:2px 5px;}
            tr:nth-child(odd){background-color:#FFE4C4;}
            tr:nth-child(even){background-color:#F0F0F0;}
            th{background-color: #0088ff;color:#FFF;}
        </style>
    </head>
    <body>
    <div class="container">
        {{ template "content" . }}
    </div>
    </body>
    </html>
{{ end }}`,

	"index.html": `{{ define "content" }}
    <h1>收支统计报告</h1>

    <table>
        <tr>
            <th rowspan="2">月份</th>
            <th colspan="2">贷款</th>
            <th colspan="6">支出</th>
            <th colspan="2">转账收支</th>
            <th>收入</th>
            <th colspan="2">其他</th>
        </tr>
        <tr>
            <th>贷款总额</th>
            <th>贷款还款</th>
            <th>总支出</th>
            <th>交通支出</th>
            <th>餐饮支出</th>
            <th>水电支出</th>
            <th>话费支出</th>
            <th>其他支出</th>
            <th>转账支出</th>
            <th>转账收入</th>
            <th>总收入</th>
            <th>信用还款</th>
            <th>内部转账</th>
        </tr>
        {{range $yearMonth := $.yearMonths}}
            {{ $monthStat := index $.monthStatsMap $yearMonth }}
        <tr>
            <td>{{$monthStat.YearMonth}}</td>
            <td>{{$monthStat.Loan.Total}}</td>
            <td>{{$monthStat.LoanRepayment.Total}}</td>
            <td>{{$monthStat.ExpenseTotal}}</td>
            <td>{{$monthStat.ExpenseTravel.Total}}</td>
            <td>{{$monthStat.ExpenseEat.Total}}</td>
            <td>{{$monthStat.ExpenseWaterElectGas.Total}}</td>
            <td>{{$monthStat.ExpenseTel.Total}}</td>
            <td>{{$monthStat.ExpenseOther.Total}}</td>
            <td>{{$monthStat.ExpenseTransfer.Total}}</td>
            <td>{{$monthStat.IncomeTransfer.Total}}</td>
            <td>{{$monthStat.Income.Total}}</td>
            <td>{{$monthStat.CreditRepayment.Total}}</td>
            <td>{{$monthStat.InnerTransfer.Total}}</td>
        </tr>
        {{end}}

    </table>

        <h2>查看明细</h2>
    {{range $yearMonth := $.yearMonths}}
        <a href="aliwepaystat-{{$yearMonth}}.html" target="stat_detail">{{$yearMonth}}</a>
    {{end}}

    <iframe name="stat_detail" style="width: 100%; height: 100%; min-height: 800px;border:0px;min-height: ">
    </iframe>
{{ end }}
`,

	"month-stat.html": `{{ define "trans_list_table"}}
    <table class="table">
        <tr>
            <th>时间</th>
            <th>交易方</th>
            <th>商品</th>
            <th>金额</th>
        </tr>
        {{range $trans := $.TransList}}
            {{if $trans.IsShowInList}}
                <tr>
                    <td>{{$trans.GetCreatedTime}}</td>
                    <td>{{$trans.GetTarget}}</td>
                    <td>{{$trans.GetProduct}}</td>
                    <td>{{$trans.GetAmount}}</td>
                </tr>
            {{end}}
        {{end}}
    </table>
{{ end }}

{{ define "content" }}
    <h1>{{$.YearMonth}} 收支统计报告</h1>

    <hr>
    <h2>1. 贷款</h2>
    <h3>1.1 贷款明细</h3>
    {{ template "trans_list_table" $.Loan }}

    <h3>1.2 贷款还款明细</h3>
    {{ template "trans_list_table" $.LoanRepayment }}

    <hr>
    <h2>2. 转账收支</h2>

    <h3>2.1 转账收入明细</h3>
    {{ template "trans_list_table" $.IncomeTransfer }}

    <h3>2.2 转账支出明细</h3>
    {{ template "trans_list_table" $.ExpenseTransfer }}

    <hr>
    <h2>3. 支出</h2>

    <h3>3.1 交通明细</h3>
    {{ template "trans_list_table" $.ExpenseTravel }}

    <h3>3.2 餐饮明细</h3>
    {{ template "trans_list_table" $.ExpenseEat }}

    <h3>3.3 水电明细</h3>
    {{ template "trans_list_table" $.ExpenseWaterElectGas }}

    <h3>3.4 话费明细</h3>
    {{ template "trans_list_table" $.ExpenseTel }}

    <h3>3.5 其他明细</h3>
    {{ template "trans_list_table" $.ExpenseOther }}

    <hr>
    <h2>4. 收入</h2>

    <h3>4.1 转账收入明细</h3>
    {{ template "trans_list_table" $.IncomeTransfer }}

    <h3>4.2 其他收入明细</h3>
    {{ template "trans_list_table" $.Income }}


    <hr>
    <h2>5. 其他</h2>
    <h3>5.1 信用还款明细</h3>
    {{ template "trans_list_table" $.CreditRepayment }}

    <h3>5.2 内部转账明细</h3>
    {{ template "trans_list_table" $.InnerTransfer }}

{{ end }}
`,
}
