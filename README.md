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

最后更新：2026-02-23 22:41:22 UTC（2026-02-24 06:41:22 UTC+8）

**代理总数：82**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 82 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 82 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 450ms | ✓ 1163ms | ✓ 1367ms | ✓ 1568ms | ✓ 1274ms | http |
| 168.235.110.63:3128 | ✓ 490ms | ✓ 1953ms | ✓ 970ms | ✓ 1287ms | ✓ 996ms | http |
| 202.152.44.19:8081 | ✓ 811ms | ✓ 1617ms | ✓ 818ms | ✓ 1200ms | ✓ 931ms | http |
| 202.152.44.18:8081 | 否 | 否 | ✓ 923ms | ✓ 1286ms | ✓ 1035ms | http |
| 211.230.49.122:3128 | ✓ 1355ms | ✓ 1014ms | ✓ 664ms | ✓ 1038ms | ✓ 802ms | http |
| 217.217.254.94:8080 | ✓ 738ms | 否 | ✓ 717ms | ✓ 1334ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1933ms | ✓ 1308ms | 否 | ✓ 1325ms | 否 | http |
| 91.107.178.131:443 | ✓ 675ms | ✓ 1861ms | ✓ 835ms | ✓ 1576ms | 否 | http |
| 34.50.41.78:8888 | ✓ 577ms | 否 | ✓ 550ms | ✓ 879ms | ✓ 694ms | http |
| 103.155.166.150:8221 | ✓ 1187ms | 否 | ✓ 1122ms | ✓ 1243ms | ✓ 1238ms | http |
| 190.242.157.215:8080 | ✓ 1704ms | ✓ 1642ms | ✓ 1038ms | ✓ 1774ms | ✓ 1516ms | http |
| 172.86.92.68:31337 | ✓ 1131ms | 否 | 否 | ✓ 1814ms | ✓ 1402ms | http |
| 120.46.152.136:3128 | ✓ 1127ms | ✓ 1340ms | ✓ 1291ms | ✓ 1180ms | 否 | http |
| 35.72.254.71:3128 | ✓ 1739ms | ✓ 1765ms | ✓ 1678ms | 否 | 否 | http |
| 137.220.150.22:6005 | ✓ 861ms | 否 | ✓ 914ms | ✓ 1096ms | ✓ 888ms | http |
| 120.92.212.16:8890 | ✓ 1017ms | ✓ 1186ms | ✓ 890ms | ✓ 1262ms | ✓ 953ms | http |
| 120.92.212.16:7890 | ✓ 1021ms | ✓ 1173ms | ✓ 891ms | 否 | ✓ 1210ms | http |
| 138.124.53.25:7443 | ✓ 717ms | 否 | 否 | ✓ 1914ms | ✓ 1520ms | http |
| 59.127.212.110:4431 | ✓ 1541ms | ✓ 1075ms | ✓ 1143ms | ✓ 1110ms | ✓ 849ms | http |
| 36.147.78.166:80 | ✓ 1615ms | 否 | ✓ 1568ms | ✓ 1932ms | 否 | http |
| 115.231.181.40:8128 | 否 | ✓ 1054ms | ✓ 897ms | 否 | ✓ 926ms | http |
| 101.43.255.96:80 | ✓ 996ms | ✓ 1225ms | ✓ 1278ms | 否 | ✓ 1008ms | http |
| 81.70.169.194:80 | ✓ 1642ms | ✓ 1321ms | ✓ 1081ms | ✓ 1232ms | ✓ 1728ms | http |
| 45.129.141.143:3128 | ✓ 824ms | 否 | ✓ 1730ms | ✓ 1977ms | ✓ 1866ms | http |
| 190.9.109.199:999 | ✓ 1038ms | ✓ 1682ms | ✓ 1330ms | ✓ 1497ms | ✓ 1594ms | http |
| 35.225.22.61:80 | ✓ 380ms | ✓ 1412ms | ✓ 553ms | ✓ 993ms | ✓ 1208ms | http |
| 103.84.95.54:7890 | ✓ 651ms | 否 | ✓ 642ms | ✓ 778ms | ✓ 675ms | http |
| 162.240.154.26:3128 | 否 | ✓ 1596ms | ✓ 420ms | ✓ 939ms | ✓ 852ms | http |
| 36.136.27.2:4999 | ✓ 1133ms | ✓ 1250ms | ✓ 1166ms | ✓ 1267ms | ✓ 1030ms | http |
| 91.238.105.64:2024 | ✓ 1112ms | ✓ 1718ms | ✓ 1307ms | ✓ 1721ms | ✓ 1359ms | http |
| 91.238.104.172:2024 | ✓ 1126ms | ✓ 1814ms | ✓ 1205ms | ✓ 1765ms | ✓ 1384ms | http |
| 85.208.108.43:10808 | ✓ 1026ms | 否 | ✓ 1358ms | ✓ 1980ms | 否 | http |
| 14.56.107.184:3128 | ✓ 1683ms | 否 | ✓ 1783ms | 否 | ✓ 1348ms | http |
| 195.123.209.48:3128 | ✓ 1273ms | ✓ 1633ms | ✓ 1791ms | ✓ 1987ms | ✓ 1863ms | http |
| 125.128.12.94:3128 | 否 | ✓ 825ms | ✓ 1426ms | ✓ 1998ms | 否 | http |
| 123.57.2.231:2020 | ✓ 1005ms | ✓ 1164ms | ✓ 936ms | ✓ 1234ms | ✓ 945ms | http |
| 35.234.17.221:8080 | ✓ 842ms | ✓ 1164ms | 否 | ✓ 1115ms | ✓ 799ms | http |
| 152.32.255.24:27197 | ✓ 1050ms | 否 | 否 | ✓ 1547ms | ✓ 1098ms | http |
| 116.80.48.38:7777 | ✓ 1605ms | 否 | ✓ 1504ms | ✓ 1793ms | ✓ 1711ms | http |
| 46.21.82.154:1080 | ✓ 596ms | 否 | ✓ 1108ms | ✓ 1637ms | ✓ 1387ms | http |
| 35.212.218.202:1080 | ✓ 1728ms | 否 | 否 | ✓ 1741ms | ✓ 1489ms | http |
| 45.140.147.82:1081 | ✓ 1702ms | ✓ 1491ms | ✓ 1471ms | 否 | 否 | http |
| 45.151.182.9:3128 | ✓ 1749ms | 否 | ✓ 1252ms | 否 | ✓ 1852ms | http |
| 8.219.97.248:80 | ✓ 1676ms | 否 | ✓ 1345ms | ✓ 1625ms | 否 | http |
| 121.204.158.249:3128 | ✓ 956ms | ✓ 1205ms | ✓ 1617ms | ✓ 1161ms | ✓ 949ms | http |
| 103.236.64.247:8888 | ✓ 878ms | 否 | ✓ 933ms | 否 | ✓ 1186ms | http |
| 103.86.131.62:80 | 否 | 否 | ✓ 1184ms | ✓ 1217ms | ✓ 975ms | http |
| 45.12.151.226:2828 | 否 | 否 | ✓ 1150ms | ✓ 1941ms | ✓ 1571ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1519ms | ✓ 1488ms | ✓ 1715ms | ✓ 1241ms | http |
| 38.180.2.107:3128 | ✓ 1732ms | ✓ 1898ms | ✓ 1852ms | 否 | 否 | http |
| 121.230.8.91:1080 | ✓ 1262ms | ✓ 1483ms | ✓ 1254ms | ✓ 1619ms | 否 | http |
| 124.16.93.70:7890 | ✓ 794ms | ✓ 1043ms | ✓ 845ms | ✓ 1055ms | ✓ 853ms | http |
| 45.140.147.82:1082 | ✓ 678ms | 否 | ✓ 1069ms | 否 | ✓ 1161ms | http |
| 91.238.104.171:2023 | 否 | ✓ 1866ms | ✓ 1686ms | 否 | ✓ 1989ms | http |
| 45.22.209.157:8888 | ✓ 1133ms | ✓ 1551ms | ✓ 794ms | 否 | 否 | http |
| 121.230.8.80:1080 | ✓ 1241ms | ✓ 1436ms | ✓ 1271ms | ✓ 1615ms | ✓ 1037ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 631ms | ✓ 950ms | ✓ 1169ms | http |
| 52.202.250.102:80 | ✓ 1687ms | ✓ 1614ms | ✓ 1521ms | 否 | ✓ 1576ms | http |
| 78.13.231.158:3128 | ✓ 1114ms | 否 | ✓ 1651ms | ✓ 1638ms | ✓ 1424ms | http |
| 165.227.5.10:8888 | 否 | ✓ 1639ms | ✓ 748ms | ✓ 1193ms | 否 | http |
| 172.212.68.37:3128 | ✓ 456ms | 否 | ✓ 1192ms | ✓ 1604ms | ✓ 1496ms | http |
| 36.147.78.166:443 | ✓ 1559ms | 否 | 否 | ✓ 1752ms | ✓ 1393ms | http |
| 103.82.23.118:5249 | 否 | 否 | ✓ 1534ms | ✓ 1763ms | ✓ 1543ms | http |
| 185.243.218.43:49153 | ✓ 1103ms | 否 | ✓ 1544ms | ✓ 1943ms | ✓ 1604ms | http |
| 103.215.36.88:10864 | 否 | ✓ 1680ms | ✓ 1593ms | 否 | ✓ 1422ms | http |
| 186.148.180.46:999 | ✓ 834ms | 否 | ✓ 868ms | ✓ 1659ms | ✓ 1496ms | http |
| 103.35.188.243:3128 | 否 | ✓ 1463ms | 否 | ✓ 1346ms | ✓ 1034ms | http |
| 178.253.22.108:65431 | ✓ 1036ms | 否 | ✓ 1436ms | ✓ 1910ms | ✓ 1202ms | http |
| 103.39.51.190:8080 | ✓ 1499ms | 否 | 否 | ✓ 1625ms | ✓ 1260ms | http |
| 202.129.206.239:3128 | ✓ 1821ms | 否 | ✓ 1795ms | ✓ 1863ms | ✓ 1291ms | http |
| 147.45.159.213:48206 | 否 | 否 | ✓ 1600ms | ✓ 1894ms | ✓ 1505ms | http |
| 18.229.170.122:3128 | ✓ 1000ms | 否 | ✓ 866ms | 否 | ✓ 1704ms | http |
| 81.177.48.54:2080 | ✓ 1209ms | 否 | ✓ 1682ms | 否 | ✓ 1605ms | http |
| 14.56.107.244:3128 | ✓ 1214ms | 否 | ✓ 984ms | 否 | ✓ 1486ms | http |
| 217.216.109.116:8080 | 否 | 否 | ✓ 985ms | ✓ 1868ms | ✓ 1051ms | http |
| 211.171.114.154:3128 | ✓ 708ms | ✓ 1849ms | 否 | ✓ 1447ms | ✓ 1548ms | http |
| 3.79.194.222:44778 | ✓ 1628ms | 否 | ✓ 1338ms | 否 | ✓ 1711ms | http |
| 18.100.127.30:3128 | ✓ 1455ms | 否 | ✓ 1572ms | 否 | ✓ 1893ms | http |
| 35.181.173.74:8841 | ✓ 1547ms | 否 | ✓ 1679ms | 否 | ✓ 1955ms | http |
| 220.197.44.36:3128 | ✓ 1990ms | 否 | ✓ 1987ms | 否 | ✓ 1773ms | http |
| 223.16.170.103:3128 | 否 | ✓ 1648ms | 否 | ✓ 1428ms | ✓ 1056ms | http |
| 103.82.23.118:5242 | 否 | ✓ 1977ms | ✓ 1277ms | 否 | ✓ 1506ms | http |

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
