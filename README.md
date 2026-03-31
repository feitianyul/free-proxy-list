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

最后更新：2026-03-31 14:41:42 UTC（2026-03-31 22:41:42 UTC+8）

**代理总数：88**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 88 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 88 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 208.87.243.199:7878 | ✓ 363ms | 否 | ✓ 1274ms | ✓ 907ms | ✓ 1714ms | http |
| 43.99.54.236:5555 | ✓ 772ms | 否 | ✓ 785ms | ✓ 958ms | ✓ 771ms | http |
| 147.161.210.140:8800 | ✓ 790ms | 否 | ✓ 1159ms | ✓ 1030ms | ✓ 1012ms | http |
| 147.161.239.240:8800 | ✓ 580ms | 否 | ✓ 1101ms | ✓ 1532ms | ✓ 1510ms | http |
| 1.231.81.166:3128 | ✓ 920ms | 否 | ✓ 1244ms | ✓ 1458ms | ✓ 1080ms | http |
| 95.213.217.168:52004 | ✓ 957ms | 否 | ✓ 1893ms | ✓ 1905ms | ✓ 1721ms | http |
| 167.103.115.102:8800 | ✓ 1665ms | 否 | ✓ 1470ms | ✓ 1421ms | ✓ 1319ms | http |
| 42.96.16.158:1311 | 否 | 否 | ✓ 1458ms | ✓ 1445ms | ✓ 1357ms | http |
| 167.103.34.108:8800 | ✓ 1945ms | 否 | ✓ 1559ms | 否 | ✓ 1992ms | http |
| 35.225.22.61:80 | ✓ 1748ms | 否 | ✓ 410ms | ✓ 1191ms | 否 | http |
| 113.160.132.26:8080 | 否 | ✓ 1605ms | ✓ 1533ms | ✓ 1461ms | 否 | http |
| 167.103.144.127:8800 | ✓ 1525ms | 否 | ✓ 1208ms | ✓ 1810ms | 否 | http |
| 150.241.71.15:1080 | 否 | ✓ 1660ms | ✓ 1628ms | ✓ 1323ms | 否 | http |
| 5.102.109.41:999 | ✓ 1172ms | 否 | ✓ 1104ms | 否 | ✓ 1354ms | http |
| 167.103.31.122:8800 | ✓ 1629ms | 否 | ✓ 1400ms | ✓ 1992ms | ✓ 1468ms | http |
| 116.80.65.75:3172 | 否 | 否 | ✓ 1636ms | ✓ 1959ms | ✓ 1797ms | http |
| 133.242.138.34:8100 | ✓ 1766ms | 否 | 否 | ✓ 1581ms | ✓ 1496ms | http |
| 209.126.84.232:8888 | ✓ 124ms | 否 | ✓ 425ms | ✓ 1249ms | ✓ 781ms | http |
| 195.19.217.200:3128 | ✓ 1593ms | 否 | ✓ 1923ms | 否 | ✓ 1895ms | http |
| 101.47.73.135:3128 | ✓ 1191ms | 否 | 否 | ✓ 1472ms | ✓ 1358ms | http |
| 177.234.217.88:999 | ✓ 1382ms | 否 | ✓ 1909ms | ✓ 1792ms | ✓ 1590ms | http |
| 20.210.39.153:8561 | 否 | ✓ 1212ms | ✓ 751ms | ✓ 1018ms | ✓ 857ms | http |
| 20.78.26.206:8561 | 否 | ✓ 1038ms | ✓ 908ms | ✓ 1039ms | ✓ 854ms | http |
| 45.136.131.51:8449 | ✓ 959ms | ✓ 850ms | ✓ 869ms | 否 | ✓ 837ms | http |
| 20.27.11.248:8561 | 否 | ✓ 1737ms | ✓ 1085ms | ✓ 1231ms | ✓ 940ms | http |
| 217.217.249.160:8080 | ✓ 1731ms | 否 | ✓ 1011ms | 否 | ✓ 1506ms | http |
| 20.78.118.91:8561 | 否 | ✓ 1260ms | ✓ 655ms | ✓ 970ms | ✓ 858ms | http |
| 20.27.15.111:8561 | 否 | ✓ 1098ms | ✓ 816ms | ✓ 1035ms | ✓ 888ms | http |
| 20.27.13.35:8561 | 否 | ✓ 1247ms | ✓ 719ms | ✓ 998ms | ✓ 881ms | http |
| 120.92.212.16:7890 | ✓ 1080ms | 否 | ✓ 1356ms | ✓ 1621ms | 否 | http |
| 20.27.14.220:8561 | 否 | ✓ 1252ms | ✓ 769ms | ✓ 995ms | ✓ 861ms | http |
| 20.210.76.104:8561 | 否 | ✓ 1261ms | ✓ 650ms | ✓ 1126ms | ✓ 851ms | http |
| 20.210.76.175:8561 | 否 | ✓ 1331ms | ✓ 638ms | ✓ 1068ms | ✓ 856ms | http |
| 20.27.15.49:8561 | 否 | ✓ 1621ms | ✓ 681ms | ✓ 1039ms | ✓ 727ms | http |
| 20.210.76.178:8561 | 否 | ✓ 1781ms | ✓ 672ms | ✓ 1078ms | ✓ 826ms | http |
| 194.59.204.87:9080 | ✓ 1129ms | ✓ 1589ms | ✓ 1465ms | 否 | ✓ 1901ms | http |
| 45.12.151.226:2829 | ✓ 1960ms | 否 | ✓ 969ms | ✓ 1564ms | ✓ 1574ms | http |
| 38.145.220.33:8448 | ✓ 831ms | 否 | ✓ 437ms | ✓ 989ms | ✓ 1183ms | http |
| 150.107.140.238:3128 | ✓ 1773ms | 否 | ✓ 1983ms | ✓ 1329ms | 否 | http |
| 38.34.179.105:8449 | 否 | 否 | ✓ 1268ms | ✓ 1007ms | ✓ 1233ms | http |
| 181.78.44.63:999 | ✓ 1035ms | 否 | ✓ 1255ms | ✓ 1538ms | ✓ 1400ms | http |
| 62.113.119.14:8080 | ✓ 1402ms | ✓ 1528ms | ✓ 997ms | ✓ 1506ms | ✓ 1190ms | http |
| 38.34.179.150:8449 | 否 | ✓ 954ms | ✓ 932ms | ✓ 1661ms | ✓ 1127ms | http |
| 38.34.179.14:8450 | 否 | ✓ 1716ms | ✓ 1181ms | 否 | ✓ 859ms | http |
| 38.145.218.161:8444 | ✓ 666ms | 否 | ✓ 247ms | ✓ 881ms | ✓ 1042ms | http |
| 137.184.1.87:3128 | 否 | ✓ 864ms | ✓ 1025ms | ✓ 890ms | ✓ 671ms | http |
| 38.34.179.19:8452 | ✓ 669ms | ✓ 1003ms | ✓ 883ms | ✓ 982ms | ✓ 746ms | http |
| 38.145.220.102:8453 | ✓ 666ms | ✓ 934ms | ✓ 731ms | ✓ 1639ms | ✓ 731ms | http |
| 38.34.179.27:8452 | ✓ 908ms | 否 | ✓ 729ms | ✓ 1187ms | ✓ 779ms | http |
| 38.34.179.50:8448 | ✓ 1192ms | ✓ 1033ms | ✓ 508ms | ✓ 1590ms | ✓ 1229ms | http |
| 38.34.183.211:8445 | ✓ 667ms | 否 | ✓ 340ms | ✓ 941ms | ✓ 1027ms | http |
| 38.34.183.16:8445 | ✓ 643ms | ✓ 941ms | ✓ 819ms | ✓ 1644ms | ✓ 811ms | http |
| 38.145.218.210:8451 | ✓ 665ms | 否 | ✓ 564ms | ✓ 995ms | ✓ 705ms | http |
| 38.145.208.186:8448 | ✓ 640ms | ✓ 869ms | ✓ 698ms | 否 | ✓ 923ms | http |
| 45.136.131.44:8449 | 否 | 否 | ✓ 738ms | ✓ 970ms | ✓ 1609ms | http |
| 38.34.183.130:8451 | 否 | 否 | ✓ 553ms | ✓ 1362ms | ✓ 1562ms | http |
| 38.34.179.66:8446 | ✓ 787ms | 否 | ✓ 1210ms | ✓ 1490ms | ✓ 955ms | http |
| 38.34.179.26:8446 | ✓ 911ms | 否 | ✓ 972ms | ✓ 1679ms | ✓ 905ms | http |
| 38.34.179.100:8452 | ✓ 1289ms | ✓ 934ms | ✓ 1421ms | 否 | ✓ 989ms | http |
| 38.34.179.84:8453 | ✓ 1183ms | 否 | ✓ 1565ms | ✓ 1848ms | ✓ 793ms | http |
| 38.145.203.135:8444 | ✓ 668ms | ✓ 1586ms | ✓ 1024ms | ✓ 1919ms | 否 | http |
| 38.34.179.193:8452 | 否 | 否 | ✓ 1747ms | ✓ 1436ms | ✓ 1175ms | http |
| 38.34.179.83:8448 | ✓ 1973ms | ✓ 1350ms | 否 | ✓ 1550ms | 否 | http |
| 128.199.114.189:9090 | ✓ 1069ms | 否 | ✓ 1865ms | ✓ 1706ms | ✓ 1683ms | http |
| 45.136.198.40:3128 | ✓ 1918ms | 否 | ✓ 1870ms | ✓ 1869ms | ✓ 1557ms | http |
| 45.136.130.197:8452 | ✓ 745ms | ✓ 1401ms | 否 | ✓ 978ms | ✓ 946ms | http |
| 120.92.212.16:8890 | ✓ 1193ms | 否 | ✓ 1313ms | ✓ 1383ms | ✓ 1118ms | http |
| 167.71.196.28:8080 | ✓ 840ms | 否 | ✓ 1677ms | ✓ 1197ms | 否 | http |
| 121.126.185.63:25152 | ✓ 1802ms | 否 | ✓ 1611ms | 否 | ✓ 1547ms | http |
| 31.192.106.135:8010 | ✓ 1473ms | 否 | ✓ 1588ms | ✓ 1796ms | ✓ 1824ms | http |
| 103.155.197.103:8080 | ✓ 1975ms | 否 | ✓ 1528ms | ✓ 1812ms | ✓ 1483ms | http |
| 115.231.181.40:8128 | ✓ 1663ms | 否 | ✓ 1421ms | 否 | ✓ 1883ms | http |
| 101.43.127.100:8877 | ✓ 1750ms | 否 | ✓ 958ms | 否 | ✓ 991ms | http |
| 38.34.179.61:8445 | ✓ 1006ms | ✓ 1271ms | ✓ 745ms | ✓ 1293ms | ✓ 1221ms | http |
| 190.12.150.244:999 | ✓ 1403ms | ✓ 1801ms | ✓ 1155ms | 否 | 否 | http |
| 109.234.38.35:3128 | ✓ 1192ms | ✓ 1623ms | ✓ 1515ms | ✓ 1581ms | ✓ 1169ms | http |
| 192.71.213.85:9091 | ✓ 833ms | 否 | ✓ 1399ms | ✓ 1707ms | 否 | http |
| 45.136.130.196:8450 | ✓ 1217ms | ✓ 1654ms | ✓ 425ms | ✓ 1036ms | ✓ 1002ms | http |
| 38.34.178.244:8446 | ✓ 1236ms | 否 | ✓ 880ms | ✓ 1627ms | ✓ 1233ms | http |
| 38.145.208.241:8453 | ✓ 407ms | ✓ 1209ms | ✓ 1108ms | ✓ 957ms | ✓ 717ms | http |
| 109.172.54.104:8888 | ✓ 1179ms | ✓ 1857ms | 否 | 否 | ✓ 1901ms | http |
| 209.126.10.139:3128 | ✓ 1423ms | 否 | 否 | ✓ 1809ms | ✓ 1309ms | http |
| 38.145.203.34:8444 | ✓ 1106ms | 否 | ✓ 790ms | ✓ 954ms | ✓ 813ms | http |
| 38.145.208.172:8448 | ✓ 332ms | ✓ 879ms | ✓ 1720ms | ✓ 902ms | ✓ 825ms | http |
| 38.145.203.132:8450 | ✓ 276ms | 否 | ✓ 1304ms | ✓ 963ms | ✓ 1063ms | http |
| 38.34.179.162:8451 | ✓ 1340ms | 否 | 否 | ✓ 966ms | ✓ 869ms | http |
| 202.129.206.239:3128 | 否 | 否 | ✓ 1787ms | ✓ 1910ms | ✓ 1464ms | http |
| 38.51.232.90:1986 | 否 | 否 | ✓ 1950ms | ✓ 1771ms | ✓ 1595ms | http |

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
