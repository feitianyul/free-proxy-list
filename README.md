**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-03-20 03:21:24 UTC（2026-03-20 11:21:24 UTC+8）

**代理总数：215**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 215 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 215 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 38.34.179.96:8451 | ✓ 1832ms | ✓ 1037ms | ✓ 542ms | ✓ 715ms | ✓ 965ms | http |
| 38.34.179.61:8445 | 否 | ✓ 1281ms | ✓ 129ms | ✓ 729ms | ✓ 896ms | http |
| 45.136.131.39:8451 | ✓ 1631ms | ✓ 908ms | ✓ 786ms | ✓ 1745ms | ✓ 625ms | http |
| 38.34.179.150:8449 | ✓ 1631ms | ✓ 1912ms | ✓ 670ms | ✓ 944ms | 否 | http |
| 103.113.70.189:1081 | ✓ 363ms | 否 | ✓ 1334ms | ✓ 1892ms | ✓ 1095ms | http |
| 38.34.179.162:8451 | ✓ 1632ms | ✓ 1921ms | ✓ 666ms | ✓ 1254ms | ✓ 1742ms | http |
| 38.34.183.13:8449 | 否 | ✓ 1608ms | ✓ 103ms | 否 | ✓ 1401ms | http |
| 38.34.179.14:8450 | ✓ 1009ms | ✓ 1587ms | ✓ 729ms | ✓ 810ms | ✓ 746ms | http |
| 38.34.179.16:8451 | ✓ 1008ms | 否 | ✓ 588ms | ✓ 670ms | ✓ 717ms | http |
| 45.136.131.54:8448 | 否 | 否 | ✓ 751ms | ✓ 1012ms | ✓ 1820ms | http |
| 147.161.210.140:8800 | 否 | 否 | ✓ 923ms | ✓ 1064ms | ✓ 947ms | http |
| 133.242.138.34:8100 | 否 | ✓ 1430ms | ✓ 1944ms | ✓ 1113ms | ✓ 1017ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1057ms | 否 | ✓ 1467ms | ✓ 845ms | http |
| 45.136.130.177:8448 | 否 | 否 | ✓ 1520ms | ✓ 830ms | ✓ 1321ms | http |
| 147.161.239.240:8800 | ✓ 765ms | ✓ 1875ms | ✓ 1401ms | ✓ 1923ms | ✓ 1722ms | http |
| 38.55.104.68:6005 | ✓ 1174ms | 否 | ✓ 1811ms | ✓ 1003ms | ✓ 1602ms | http |
| 38.34.179.27:8451 | 否 | ✓ 1456ms | ✓ 1188ms | ✓ 1867ms | ✓ 654ms | http |
| 113.160.132.26:8080 | ✓ 1318ms | 否 | ✓ 1276ms | 否 | ✓ 1399ms | http |
| 174.138.24.77:1080 | 否 | 否 | ✓ 1407ms | ✓ 1171ms | ✓ 877ms | http |
| 45.136.130.171:8445 | ✓ 285ms | ✓ 804ms | ✓ 128ms | ✓ 744ms | ✓ 1126ms | http |
| 38.34.183.234:8450 | ✓ 323ms | ✓ 671ms | ✓ 209ms | ✓ 924ms | ✓ 1662ms | http |
| 38.145.218.228:8447 | 否 | ✓ 613ms | ✓ 197ms | ✓ 921ms | 否 | http |
| 38.34.179.178:8445 | ✓ 1751ms | ✓ 742ms | ✓ 640ms | 否 | ✓ 514ms | http |
| 38.34.179.97:8448 | ✓ 576ms | ✓ 1367ms | ✓ 654ms | ✓ 665ms | ✓ 583ms | http |
| 45.136.131.35:8452 | ✓ 1718ms | ✓ 1687ms | ✓ 644ms | ✓ 1110ms | ✓ 1071ms | http |
| 38.34.183.8:8450 | ✓ 321ms | ✓ 1193ms | ✓ 1485ms | ✓ 649ms | ✓ 550ms | http |
| 38.34.179.60:8450 | ✓ 304ms | ✓ 1210ms | ✓ 1814ms | ✓ 748ms | ✓ 512ms | http |
| 38.145.220.33:8448 | 否 | ✓ 1322ms | ✓ 641ms | 否 | ✓ 713ms | http |
| 38.34.179.105:8449 | ✓ 1782ms | ✓ 1716ms | ✓ 1537ms | ✓ 751ms | ✓ 544ms | http |
| 219.117.204.211:7799 | ✓ 697ms | 否 | 否 | ✓ 968ms | ✓ 940ms | http |
| 38.34.179.184:8450 | ✓ 321ms | ✓ 1322ms | 否 | ✓ 1088ms | ✓ 737ms | http |
| 38.34.179.190:8450 | ✓ 326ms | ✓ 1107ms | 否 | ✓ 1309ms | ✓ 744ms | http |
| 38.34.179.192:8450 | ✓ 294ms | ✓ 958ms | 否 | ✓ 1461ms | ✓ 745ms | http |
| 137.220.150.152:6005 | 否 | 否 | ✓ 755ms | ✓ 1353ms | ✓ 924ms | http |
| 45.125.67.37:443 | 否 | 否 | ✓ 1438ms | ✓ 1027ms | ✓ 1040ms | http |
| 137.220.150.22:6005 | 否 | ✓ 1868ms | ✓ 1470ms | ✓ 1282ms | ✓ 814ms | http |
| 45.136.130.186:8451 | ✓ 1766ms | ✓ 1165ms | ✓ 1221ms | ✓ 1675ms | ✓ 571ms | http |
| 45.136.131.36:8450 | ✓ 1742ms | ✓ 654ms | ✓ 224ms | ✓ 1316ms | 否 | http |
| 38.145.220.198:8448 | 否 | ✓ 839ms | ✓ 1279ms | 否 | ✓ 723ms | http |
| 193.23.200.251:10808 | ✓ 1175ms | 否 | ✓ 1680ms | 否 | ✓ 1656ms | http |
| 38.55.104.99:6005 | 否 | 否 | ✓ 1266ms | ✓ 1401ms | ✓ 1504ms | http |
| 38.34.178.154:8445 | 否 | ✓ 1468ms | 否 | ✓ 1306ms | ✓ 601ms | http |
| 45.167.124.52:8080 | ✓ 888ms | ✓ 1549ms | ✓ 704ms | ✓ 1744ms | ✓ 1535ms | http |
| 59.46.216.131:30001 | ✓ 1044ms | 否 | 否 | ✓ 1948ms | ✓ 1122ms | http |
| 164.90.155.209:3128 | ✓ 292ms | ✓ 1176ms | ✓ 929ms | ✓ 631ms | ✓ 492ms | http |
| 137.184.1.87:3128 | ✓ 572ms | ✓ 1876ms | ✓ 603ms | ✓ 800ms | ✓ 500ms | http |
| 38.145.220.11:8445 | ✓ 439ms | ✓ 604ms | ✓ 410ms | ✓ 800ms | ✓ 1219ms | http |
| 38.145.218.163:8451 | ✓ 704ms | 否 | ✓ 599ms | ✓ 670ms | ✓ 516ms | http |
| 20.210.39.153:8561 | ✓ 1232ms | ✓ 1099ms | ✓ 668ms | ✓ 909ms | ✓ 574ms | http |
| 20.78.26.206:8561 | ✓ 1233ms | ✓ 1616ms | ✓ 435ms | ✓ 741ms | ✓ 583ms | http |
| 20.78.118.91:8561 | ✓ 1245ms | ✓ 1691ms | ✓ 437ms | ✓ 768ms | ✓ 579ms | http |
| 38.34.179.63:8448 | ✓ 1860ms | ✓ 1036ms | ✓ 185ms | ✓ 715ms | ✓ 1631ms | http |
| 194.59.204.87:9080 | ✓ 645ms | ✓ 1724ms | ✓ 1720ms | 否 | 否 | http |
| 46.101.190.71:3128 | ✓ 633ms | 否 | ✓ 1891ms | ✓ 1682ms | ✓ 1323ms | http |
| 115.231.181.40:8128 | ✓ 931ms | 否 | ✓ 935ms | ✓ 1380ms | ✓ 1971ms | http |
| 154.64.243.50:7890 | ✓ 1518ms | ✓ 1972ms | ✓ 1455ms | 否 | ✓ 1654ms | http |
| 38.34.179.49:8450 | 否 | ✓ 1388ms | ✓ 184ms | ✓ 821ms | 否 | http |
| 91.238.104.171:2023 | ✓ 1263ms | 否 | ✓ 1692ms | 否 | ✓ 1333ms | http |
| 222.228.171.92:8080 | ✓ 1984ms | 否 | ✓ 1749ms | ✓ 808ms | 否 | http |
| 35.225.22.61:80 | ✓ 586ms | ✓ 1572ms | ✓ 308ms | ✓ 1154ms | ✓ 1034ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1207ms | ✓ 926ms | ✓ 1211ms | ✓ 936ms | http |
| 167.103.34.108:8800 | ✓ 1239ms | 否 | ✓ 1476ms | ✓ 1296ms | ✓ 1225ms | http |
| 137.220.150.104:6005 | ✓ 792ms | ✓ 1997ms | ✓ 888ms | ✓ 1612ms | 否 | http |
| 167.103.31.122:8800 | ✓ 1285ms | 否 | ✓ 1362ms | 否 | ✓ 1476ms | http |
| 38.145.220.32:8444 | ✓ 441ms | ✓ 668ms | ✓ 969ms | ✓ 992ms | ✓ 509ms | http |
| 45.136.131.46:8445 | ✓ 684ms | ✓ 953ms | ✓ 446ms | ✓ 967ms | ✓ 516ms | http |
| 101.43.127.100:8877 | 否 | ✓ 965ms | ✓ 1454ms | ✓ 1088ms | ✓ 846ms | http |
| 38.145.208.169:8450 | ✓ 803ms | ✓ 800ms | ✓ 785ms | 否 | ✓ 1031ms | http |
| 38.145.208.226:8451 | ✓ 1195ms | ✓ 696ms | 否 | ✓ 1912ms | ✓ 1689ms | http |
| 38.145.220.96:8448 | ✓ 1875ms | ✓ 604ms | ✓ 407ms | ✓ 1839ms | ✓ 1678ms | http |
| 210.223.44.230:3128 | ✓ 677ms | ✓ 933ms | ✓ 1080ms | ✓ 1080ms | ✓ 709ms | http |
| 137.220.151.110:6005 | ✓ 989ms | 否 | ✓ 1129ms | ✓ 1331ms | ✓ 810ms | http |
| 101.47.73.135:3128 | ✓ 1459ms | 否 | ✓ 1187ms | ✓ 1302ms | 否 | http |
| 85.208.108.43:2094 | ✓ 1393ms | 否 | ✓ 530ms | ✓ 1404ms | ✓ 787ms | http |
| 116.80.65.81:3172 | ✓ 1585ms | 否 | ✓ 1458ms | ✓ 1800ms | 否 | http |
| 38.34.179.176:8444 | ✓ 92ms | ✓ 612ms | ✓ 106ms | ✓ 690ms | ✓ 523ms | http |
| 167.71.60.190:8080 | ✓ 1501ms | 否 | ✓ 1634ms | ✓ 1776ms | 否 | http |
| 24.199.124.151:3128 | ✓ 755ms | ✓ 1414ms | ✓ 710ms | ✓ 673ms | ✓ 484ms | http |
| 165.227.5.10:8888 | ✓ 143ms | 否 | ✓ 967ms | ✓ 1719ms | 否 | http |
| 120.92.212.16:7890 | ✓ 917ms | ✓ 1149ms | 否 | 否 | ✓ 967ms | http |
| 211.171.114.154:3128 | 否 | 否 | ✓ 1754ms | ✓ 1369ms | ✓ 1124ms | http |
| 5.104.87.17:8050 | ✓ 1401ms | 否 | ✓ 1232ms | ✓ 1888ms | 否 | http |
| 120.132.97.88:7897 | 否 | 否 | ✓ 797ms | ✓ 1375ms | ✓ 859ms | http |
| 186.96.111.214:999 | ✓ 821ms | ✓ 1542ms | ✓ 685ms | ✓ 1548ms | ✓ 1320ms | http |
| 38.145.208.181:8445 | ✓ 162ms | ✓ 697ms | ✓ 139ms | ✓ 709ms | ✓ 610ms | http |
| 157.230.220.25:4857 | ✓ 393ms | 否 | ✓ 1886ms | ✓ 1991ms | 否 | http |
| 38.145.218.229:8450 | ✓ 121ms | ✓ 828ms | ✓ 535ms | ✓ 1150ms | ✓ 515ms | http |
| 38.34.179.20:8445 | ✓ 171ms | ✓ 619ms | ✓ 518ms | ✓ 1251ms | ✓ 1280ms | http |
| 38.34.183.224:8448 | ✓ 357ms | ✓ 646ms | ✓ 143ms | ✓ 713ms | ✓ 638ms | http |
| 38.34.183.225:8450 | ✓ 425ms | ✓ 745ms | ✓ 1826ms | 否 | ✓ 803ms | http |
| 45.136.130.190:8447 | ✓ 913ms | ✓ 678ms | ✓ 446ms | 否 | 否 | http |
| 45.136.131.58:8450 | ✓ 925ms | ✓ 829ms | ✓ 766ms | ✓ 657ms | ✓ 599ms | http |
| 64.31.49.174:3128 | ✓ 1077ms | 否 | 否 | ✓ 1614ms | ✓ 1490ms | http |
| 194.67.99.223:1080 | ✓ 1423ms | 否 | ✓ 1631ms | ✓ 1943ms | 否 | http |
| 103.84.95.54:7890 | ✓ 1848ms | 否 | ✓ 1859ms | 否 | ✓ 1778ms | http |
| 137.220.150.170:6005 | 否 | 否 | ✓ 938ms | ✓ 1558ms | ✓ 867ms | http |
| 175.194.173.105:3128 | 否 | ✓ 1506ms | ✓ 870ms | ✓ 939ms | 否 | http |
| 112.163.160.93:3128 | 否 | ✓ 1344ms | ✓ 834ms | ✓ 942ms | 否 | http |
| 109.120.185.119:8118 | ✓ 720ms | 否 | ✓ 1719ms | 否 | ✓ 1744ms | http |
| 45.136.130.193:8447 | 否 | ✓ 1159ms | ✓ 1045ms | ✓ 684ms | ✓ 707ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
