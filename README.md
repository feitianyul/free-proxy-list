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

最后更新：2026-03-29 16:29:35 UTC（2026-03-30 00:29:35 UTC+8）

**代理总数：87**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 87 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 87 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 39.185.46.193:5911 | ✓ 653ms | ✓ 794ms | ✓ 651ms | ✓ 846ms | ✓ 666ms | http |
| 147.161.210.140:8800 | ✓ 1917ms | 否 | ✓ 1105ms | ✓ 1988ms | ✓ 961ms | http |
| 113.160.132.26:8080 | ✓ 1976ms | ✓ 1394ms | 否 | ✓ 1305ms | ✓ 991ms | http |
| 42.96.16.158:1311 | ✓ 1699ms | 否 | ✓ 1448ms | ✓ 1223ms | ✓ 1131ms | http |
| 167.103.115.102:8800 | 否 | 否 | ✓ 1528ms | ✓ 1065ms | ✓ 1295ms | http |
| 167.103.34.108:8800 | 否 | 否 | ✓ 1655ms | ✓ 1333ms | ✓ 1864ms | http |
| 20.210.39.153:8561 | ✓ 1920ms | ✓ 1576ms | ✓ 1377ms | ✓ 1874ms | 否 | http |
| 20.78.26.206:8561 | ✓ 1919ms | ✓ 1511ms | ✓ 1593ms | ✓ 1931ms | 否 | http |
| 8.219.97.248:80 | ✓ 1051ms | 否 | ✓ 1231ms | ✓ 1615ms | 否 | http |
| 208.87.243.199:7878 | ✓ 223ms | ✓ 705ms | ✓ 174ms | ✓ 689ms | ✓ 552ms | http |
| 167.103.144.127:8800 | ✓ 1595ms | 否 | ✓ 753ms | ✓ 1234ms | ✓ 1090ms | http |
| 95.213.217.168:52004 | ✓ 893ms | 否 | ✓ 698ms | ✓ 1755ms | ✓ 1263ms | http |
| 167.103.31.122:8800 | ✓ 1626ms | 否 | ✓ 1333ms | 否 | ✓ 1526ms | http |
| 198.59.68.130:3128 | ✓ 713ms | ✓ 1285ms | ✓ 1981ms | ✓ 1876ms | ✓ 1088ms | http |
| 103.84.95.54:7890 | ✓ 978ms | 否 | ✓ 709ms | ✓ 1062ms | 否 | http |
| 43.99.54.236:5555 | ✓ 765ms | ✓ 1160ms | ✓ 944ms | ✓ 812ms | ✓ 649ms | http |
| 219.117.204.211:7799 | ✓ 1444ms | 否 | ✓ 594ms | ✓ 935ms | ✓ 647ms | http |
| 45.12.151.226:2829 | ✓ 800ms | 否 | ✓ 1648ms | 否 | ✓ 1473ms | http |
| 120.92.212.16:7890 | ✓ 1255ms | ✓ 1210ms | 否 | ✓ 1880ms | 否 | http |
| 1.231.81.166:3128 | ✓ 1603ms | 否 | 否 | ✓ 952ms | ✓ 818ms | http |
| 120.92.212.16:8890 | ✓ 1047ms | ✓ 1242ms | ✓ 1260ms | 否 | ✓ 954ms | http |
| 35.225.22.61:80 | ✓ 825ms | ✓ 1850ms | ✓ 1178ms | ✓ 1146ms | 否 | http |
| 147.161.239.240:8800 | ✓ 1194ms | 否 | ✓ 1100ms | ✓ 1436ms | ✓ 1611ms | http |
| 45.144.28.81:10808 | ✓ 754ms | 否 | ✓ 1923ms | ✓ 1488ms | ✓ 1312ms | http |
| 38.145.203.109:8450 | ✓ 1229ms | ✓ 1519ms | ✓ 155ms | ✓ 931ms | ✓ 704ms | http |
| 38.145.208.213:8449 | 否 | ✓ 1182ms | ✓ 994ms | 否 | ✓ 1100ms | http |
| 222.228.171.92:8080 | 否 | ✓ 1766ms | 否 | ✓ 1012ms | ✓ 1081ms | http |
| 59.46.216.131:30001 | ✓ 1781ms | ✓ 1347ms | ✓ 1457ms | 否 | 否 | http |
| 193.233.22.29:10808 | ✓ 615ms | 否 | ✓ 317ms | ✓ 1488ms | ✓ 1588ms | http |
| 91.238.123.111:8000 | ✓ 1049ms | 否 | ✓ 1034ms | 否 | ✓ 1338ms | http |
| 101.43.127.100:8877 | 否 | ✓ 1581ms | ✓ 961ms | ✓ 1428ms | ✓ 1641ms | http |
| 5.104.87.17:8051 | 否 | 否 | ✓ 1913ms | ✓ 1721ms | ✓ 1354ms | http |
| 177.234.217.88:999 | ✓ 1294ms | 否 | ✓ 1835ms | ✓ 1773ms | ✓ 1522ms | http |
| 103.52.115.171:3128 | ✓ 1784ms | 否 | ✓ 1140ms | ✓ 1210ms | ✓ 952ms | http |
| 45.149.92.147:5001 | ✓ 1045ms | 否 | ✓ 674ms | ✓ 1068ms | ✓ 824ms | http |
| 103.82.23.118:5234 | ✓ 1440ms | 否 | ✓ 1086ms | ✓ 1956ms | ✓ 1369ms | http |
| 20.78.118.91:8561 | ✓ 1976ms | ✓ 1838ms | ✓ 1633ms | 否 | 否 | http |
| 24.144.86.173:1080 | ✓ 293ms | ✓ 869ms | ✓ 895ms | ✓ 705ms | ✓ 521ms | http |
| 45.8.157.38:3128 | ✓ 696ms | 否 | ✓ 1279ms | ✓ 1896ms | ✓ 1253ms | http |
| 150.107.140.238:3128 | ✓ 944ms | 否 | ✓ 845ms | ✓ 1228ms | 否 | http |
| 121.126.185.63:25152 | 否 | ✓ 1744ms | ✓ 1283ms | ✓ 1797ms | 否 | http |
| 38.145.203.35:8450 | ✓ 1179ms | 否 | ✓ 173ms | ✓ 979ms | ✓ 1215ms | http |
| 103.100.159.145:80 | ✓ 676ms | ✓ 1491ms | ✓ 851ms | ✓ 814ms | ✓ 663ms | http |
| 164.90.206.15:3128 | ✓ 621ms | 否 | ✓ 985ms | ✓ 1450ms | ✓ 1334ms | http |
| 101.47.73.135:3128 | ✓ 1032ms | 否 | ✓ 1459ms | ✓ 1566ms | 否 | http |
| 38.34.179.67:8451 | ✓ 1278ms | ✓ 1471ms | ✓ 117ms | ✓ 719ms | ✓ 1697ms | http |
| 51.79.207.21:8080 | ✓ 1717ms | 否 | ✓ 1834ms | ✓ 1749ms | ✓ 1235ms | http |
| 5.102.109.41:999 | ✓ 1000ms | 否 | ✓ 812ms | ✓ 1657ms | ✓ 1315ms | http |
| 38.34.179.29:8449 | ✓ 666ms | ✓ 1374ms | ✓ 388ms | ✓ 672ms | ✓ 512ms | http |
| 38.145.218.227:8451 | ✓ 662ms | 否 | ✓ 1164ms | ✓ 657ms | ✓ 590ms | http |
| 38.145.208.171:8453 | ✓ 1562ms | ✓ 1521ms | ✓ 853ms | ✓ 660ms | ✓ 540ms | http |
| 38.34.179.202:8449 | ✓ 937ms | ✓ 1899ms | ✓ 348ms | ✓ 851ms | ✓ 1616ms | http |
| 38.145.208.169:8446 | ✓ 664ms | ✓ 914ms | 否 | ✓ 917ms | ✓ 611ms | http |
| 38.34.179.66:8446 | ✓ 666ms | ✓ 1543ms | ✓ 801ms | ✓ 741ms | ✓ 585ms | http |
| 38.34.179.62:8453 | ✓ 1670ms | ✓ 1908ms | ✓ 143ms | ✓ 810ms | ✓ 1594ms | http |
| 38.145.208.185:8449 | 否 | ✓ 1116ms | ✓ 638ms | ✓ 1675ms | ✓ 1153ms | http |
| 38.34.179.100:8452 | ✓ 679ms | ✓ 1342ms | ✓ 1602ms | ✓ 644ms | ✓ 950ms | http |
| 38.145.218.228:8452 | 否 | ✓ 1365ms | ✓ 983ms | 否 | ✓ 514ms | http |
| 38.34.183.13:8445 | ✓ 662ms | ✓ 1270ms | ✓ 1313ms | ✓ 670ms | ✓ 554ms | http |
| 38.145.220.198:8451 | 否 | ✓ 1384ms | ✓ 945ms | 否 | ✓ 571ms | http |
| 38.145.208.204:8445 | ✓ 677ms | ✓ 1654ms | ✓ 1760ms | ✓ 676ms | ✓ 673ms | http |
| 38.34.179.20:8452 | ✓ 679ms | ✓ 946ms | 否 | ✓ 1349ms | ✓ 911ms | http |
| 38.34.179.12:8452 | ✓ 677ms | ✓ 921ms | 否 | ✓ 1379ms | ✓ 943ms | http |
| 38.145.208.193:8451 | ✓ 662ms | ✓ 1200ms | 否 | ✓ 1199ms | ✓ 842ms | http |
| 45.136.131.38:8445 | ✓ 1652ms | ✓ 766ms | ✓ 1896ms | 否 | ✓ 545ms | http |
| 195.123.209.48:3128 | ✓ 1445ms | ✓ 1664ms | ✓ 1820ms | 否 | ✓ 1868ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1974ms | ✓ 1219ms | ✓ 1336ms | http |
| 106.117.208.101:7890 | 否 | 否 | ✓ 1028ms | ✓ 1236ms | ✓ 1752ms | http |
| 181.79.95.65:999 | ✓ 1472ms | ✓ 1780ms | ✓ 1598ms | ✓ 1947ms | 否 | http |
| 45.144.232.5:11741 | ✓ 1839ms | 否 | ✓ 1651ms | 否 | ✓ 1966ms | http |
| 181.41.201.85:3128 | ✓ 1841ms | 否 | ✓ 1996ms | 否 | ✓ 1975ms | http |
| 116.80.65.75:3172 | ✓ 1946ms | 否 | ✓ 1865ms | 否 | ✓ 1657ms | http |
| 45.136.198.40:3128 | ✓ 797ms | 否 | ✓ 1621ms | 否 | ✓ 1863ms | http |
| 103.39.51.190:8080 | ✓ 1784ms | 否 | 否 | ✓ 1504ms | ✓ 1441ms | http |
| 38.145.208.235:8445 | ✓ 979ms | ✓ 623ms | ✓ 415ms | ✓ 718ms | ✓ 1498ms | http |
| 152.69.229.220:3128 | ✓ 1478ms | ✓ 1722ms | ✓ 1222ms | ✓ 1074ms | ✓ 935ms | http |
| 59.8.203.55:80 | ✓ 1475ms | ✓ 1153ms | ✓ 957ms | ✓ 951ms | 否 | http |
| 143.244.140.119:3128 | ✓ 1345ms | 否 | ✓ 1644ms | ✓ 1760ms | ✓ 1603ms | http |
| 223.16.170.103:80 | ✓ 1451ms | 否 | 否 | ✓ 1235ms | ✓ 910ms | http |
| 103.183.10.169:3125 | ✓ 1782ms | 否 | 否 | ✓ 1672ms | ✓ 1480ms | http |
| 144.124.227.90:21074 | ✓ 1211ms | 否 | ✓ 1976ms | 否 | ✓ 1682ms | http |
| 201.150.116.3:999 | 否 | ✓ 1447ms | 否 | ✓ 1556ms | ✓ 1343ms | http |
| 114.237.77.239:1080 | ✓ 1000ms | ✓ 1938ms | ✓ 977ms | ✓ 1239ms | ✓ 1705ms | http |
| 45.136.130.166:8449 | ✓ 119ms | ✓ 742ms | ✓ 998ms | ✓ 690ms | ✓ 541ms | http |
| 45.136.131.30:8446 | ✓ 216ms | 否 | ✓ 157ms | ✓ 719ms | ✓ 992ms | http |
| 38.34.179.35:8448 | ✓ 447ms | 否 | ✓ 818ms | ✓ 768ms | ✓ 731ms | http |
| 113.255.59.226:8080 | 否 | 否 | ✓ 1213ms | ✓ 1106ms | ✓ 1100ms | http |

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
