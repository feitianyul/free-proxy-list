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

最后更新：2026-03-17 19:58:40 UTC（2026-03-18 03:58:40 UTC+8）

**代理总数：75**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 74 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 75 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 5.129.202.104:8888 | ✓ 956ms | ✓ 1956ms | ✓ 1726ms | 否 | 否 | http |
| 202.155.12.161:443 | ✓ 1468ms | 否 | ✓ 1422ms | ✓ 1523ms | 否 | http |
| 147.161.239.240:8800 | ✓ 870ms | ✓ 1487ms | ✓ 1630ms | ✓ 1581ms | ✓ 1253ms | http |
| 147.161.210.140:8800 | ✓ 1466ms | ✓ 1522ms | ✓ 1050ms | ✓ 1372ms | ✓ 1310ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1526ms | ✓ 1385ms | ✓ 1642ms | ✓ 1105ms | http |
| 45.167.124.52:8080 | ✓ 1053ms | 否 | ✓ 1482ms | 否 | ✓ 1366ms | http |
| 219.117.204.211:7799 | ✓ 995ms | 否 | ✓ 716ms | 否 | ✓ 1171ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1189ms | ✓ 1446ms | ✓ 1117ms | http |
| 1.231.81.166:3128 | ✓ 867ms | ✓ 1643ms | ✓ 1348ms | ✓ 1458ms | ✓ 988ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1680ms | 否 | ✓ 1690ms | ✓ 1427ms | http |
| 138.124.53.25:7443 | ✓ 1300ms | 否 | ✓ 990ms | 否 | ✓ 1372ms | http |
| 137.220.150.152:6005 | ✓ 1930ms | 否 | ✓ 929ms | ✓ 1451ms | ✓ 1140ms | http |
| 190.12.150.244:999 | ✓ 1947ms | 否 | ✓ 1853ms | 否 | ✓ 1752ms | http |
| 37.187.109.70:10111 | ✓ 1203ms | ✓ 1188ms | ✓ 927ms | ✓ 1476ms | ✓ 1614ms | http |
| 137.220.150.170:6005 | 否 | 否 | ✓ 1156ms | ✓ 1638ms | ✓ 1340ms | http |
| 59.46.216.131:30001 | ✓ 1109ms | ✓ 1537ms | ✓ 1380ms | 否 | ✓ 1276ms | http |
| 101.43.127.100:8877 | ✓ 1018ms | ✓ 1351ms | ✓ 1010ms | ✓ 1357ms | ✓ 996ms | http |
| 133.242.138.34:8100 | 否 | ✓ 1867ms | ✓ 1280ms | ✓ 1806ms | 否 | http |
| 85.198.96.242:3128 | ✓ 503ms | 否 | ✓ 537ms | ✓ 1997ms | 否 | http |
| 45.186.6.104:3128 | ✓ 1327ms | ✓ 1823ms | ✓ 1524ms | 否 | 否 | http |
| 160.250.4.245:1 | 否 | 否 | ✓ 1457ms | ✓ 1770ms | ✓ 1402ms | http |
| 94.72.109.169:8080 | ✓ 423ms | 否 | ✓ 1826ms | ✓ 1493ms | ✓ 1509ms | http |
| 113.176.92.71:3128 | ✓ 1426ms | ✓ 1690ms | ✓ 1413ms | ✓ 1750ms | ✓ 1468ms | http |
| 185.115.74.185:8080 | ✓ 905ms | ✓ 1546ms | ✓ 1595ms | 否 | 否 | http |
| 65.108.203.36:18080 | ✓ 983ms | 否 | 否 | ✓ 1993ms | ✓ 1784ms | http |
| 150.249.255.91:3128 | ✓ 1644ms | 否 | ✓ 692ms | 否 | ✓ 1660ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1271ms | ✓ 1298ms | ✓ 1319ms | http |
| 88.80.150.82:8080 | ✓ 970ms | 否 | ✓ 1531ms | ✓ 1972ms | ✓ 1216ms | https |
| 137.220.150.104:6005 | ✓ 973ms | 否 | ✓ 1185ms | ✓ 1431ms | ✓ 1518ms | http |
| 194.5.212.40:8080 | ✓ 429ms | 否 | ✓ 1580ms | ✓ 1922ms | ✓ 1897ms | http |
| 137.220.151.110:6005 | ✓ 970ms | 否 | ✓ 1148ms | ✓ 1462ms | ✓ 1472ms | http |
| 168.235.110.63:3128 | ✓ 606ms | ✓ 907ms | ✓ 351ms | ✓ 1529ms | ✓ 1941ms | http |
| 178.236.245.59:3128 | ✓ 948ms | 否 | ✓ 798ms | ✓ 1652ms | 否 | http |
| 178.236.245.17:3128 | ✓ 954ms | 否 | ✓ 798ms | ✓ 1694ms | 否 | http |
| 38.145.208.239:8443 | ✓ 748ms | ✓ 1413ms | ✓ 443ms | ✓ 1118ms | ✓ 684ms | http |
| 38.145.218.228:8447 | ✓ 756ms | ✓ 933ms | ✓ 928ms | ✓ 1079ms | ✓ 900ms | http |
| 38.145.208.166:8443 | ✓ 755ms | ✓ 985ms | ✓ 881ms | ✓ 1082ms | ✓ 860ms | http |
| 38.145.208.160:8443 | ✓ 757ms | ✓ 925ms | ✓ 942ms | ✓ 1074ms | ✓ 882ms | http |
| 38.145.208.164:8443 | ✓ 742ms | ✓ 965ms | ✓ 897ms | ✓ 1096ms | ✓ 870ms | http |
| 38.145.208.168:8443 | ✓ 750ms | ✓ 979ms | ✓ 889ms | ✓ 1083ms | ✓ 862ms | http |
| 38.145.208.167:8443 | ✓ 745ms | ✓ 1007ms | ✓ 849ms | ✓ 1070ms | ✓ 908ms | http |
| 38.145.208.165:8443 | ✓ 767ms | ✓ 1543ms | ✓ 466ms | ✓ 962ms | ✓ 860ms | http |
| 38.145.208.151:8443 | ✓ 769ms | ✓ 994ms | ✓ 867ms | ✓ 1118ms | ✓ 876ms | http |
| 38.145.208.163:8443 | ✓ 756ms | ✓ 965ms | ✓ 894ms | ✓ 1078ms | ✓ 902ms | http |
| 59.8.203.55:80 | ✓ 1399ms | ✓ 1319ms | ✓ 1190ms | ✓ 1189ms | ✓ 976ms | http |
| 148.135.116.20:8118 | ✓ 818ms | ✓ 873ms | ✓ 1056ms | 否 | 否 | http |
| 106.14.203.63:3333 | 否 | ✓ 1268ms | ✓ 1010ms | 否 | ✓ 1059ms | http |
| 14.143.222.113:10155 | ✓ 1598ms | 否 | ✓ 975ms | ✓ 1343ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1993ms | 否 | 否 | ✓ 1546ms | ✓ 1572ms | http |
| 165.227.5.10:8888 | ✓ 843ms | ✓ 1236ms | 否 | 否 | ✓ 865ms | http |
| 38.145.208.177:8443 | ✓ 843ms | ✓ 834ms | ✓ 330ms | ✓ 868ms | ✓ 783ms | http |
| 38.145.208.174:8443 | ✓ 857ms | ✓ 886ms | ✓ 284ms | ✓ 960ms | ✓ 748ms | http |
| 38.145.218.227:8445 | ✓ 711ms | ✓ 1015ms | ✓ 312ms | ✓ 930ms | ✓ 718ms | http |
| 38.145.218.51:8443 | ✓ 318ms | ✓ 875ms | ✓ 437ms | ✓ 952ms | ✓ 731ms | http |
| 38.145.218.13:8443 | ✓ 326ms | ✓ 878ms | ✓ 433ms | ✓ 966ms | ✓ 727ms | http |
| 45.136.130.252:8443 | ✓ 303ms | ✓ 832ms | ✓ 282ms | ✓ 936ms | ✓ 705ms | http |
| 45.136.130.247:8443 | ✓ 308ms | ✓ 879ms | ✓ 308ms | ✓ 920ms | ✓ 712ms | http |
| 45.136.131.41:8451 | 否 | ✓ 1856ms | ✓ 335ms | ✓ 1148ms | ✓ 1121ms | http |
| 45.136.131.36:8450 | ✓ 556ms | ✓ 1574ms | ✓ 345ms | ✓ 1459ms | ✓ 1320ms | http |
| 38.34.179.71:8447 | ✓ 794ms | ✓ 1763ms | ✓ 1063ms | 否 | ✓ 813ms | http |
| 45.136.131.64:8445 | 否 | ✓ 972ms | ✓ 506ms | ✓ 976ms | ✓ 712ms | http |
| 192.3.203.158:1080 | ✓ 1371ms | ✓ 1770ms | ✓ 794ms | ✓ 1701ms | ✓ 1333ms | http |
| 38.34.179.21:8446 | 否 | 否 | ✓ 1552ms | ✓ 940ms | ✓ 803ms | http |
| 45.136.131.59:8450 | 否 | 否 | ✓ 901ms | ✓ 1255ms | ✓ 738ms | http |
| 213.3.34.39:443 | ✓ 1042ms | ✓ 1955ms | ✓ 1256ms | ✓ 1866ms | ✓ 1155ms | http |
| 62.234.206.73:3128 | ✓ 1175ms | ✓ 1457ms | ✓ 1110ms | ✓ 1525ms | ✓ 1155ms | http |
| 222.109.119.178:3128 | ✓ 1731ms | ✓ 1683ms | ✓ 1282ms | 否 | 否 | http |
| 45.136.130.187:8452 | ✓ 955ms | ✓ 1475ms | ✓ 424ms | ✓ 1057ms | ✓ 811ms | http |
| 38.145.208.160:8453 | ✓ 945ms | ✓ 887ms | ✓ 1166ms | ✓ 1724ms | ✓ 739ms | http |
| 38.145.220.65:8448 | ✓ 958ms | ✓ 884ms | ✓ 997ms | ✓ 1193ms | ✓ 968ms | http |
| 38.145.220.39:8447 | ✓ 1336ms | ✓ 1024ms | ✓ 1907ms | ✓ 1600ms | ✓ 858ms | http |
| 144.31.137.23:8080 | ✓ 1460ms | 否 | ✓ 750ms | 否 | ✓ 1676ms | http |
| 106.117.208.101:7890 | ✓ 1204ms | ✓ 1406ms | ✓ 1275ms | ✓ 1531ms | ✓ 1270ms | http |
| 38.145.218.14:8443 | ✓ 395ms | ✓ 1434ms | ✓ 1077ms | ✓ 944ms | ✓ 690ms | http |
| 85.8.182.108:443 | ✓ 1055ms | 否 | ✓ 1144ms | ✓ 1361ms | ✓ 839ms | http |

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
