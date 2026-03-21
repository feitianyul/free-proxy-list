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

最后更新：2026-03-21 18:19:48 UTC（2026-03-22 02:19:48 UTC+8）

**代理总数：73**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 72 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 73 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1494ms | 否 | ✓ 1163ms | ✓ 1226ms | ✓ 769ms | http |
| 147.161.239.240:8800 | ✓ 1242ms | 否 | ✓ 1513ms | ✓ 1824ms | ✓ 1196ms | http |
| 113.160.132.26:8080 | ✓ 1490ms | 否 | ✓ 1698ms | ✓ 1405ms | ✓ 1835ms | http |
| 45.167.124.52:8080 | ✓ 1985ms | ✓ 1968ms | ✓ 1360ms | ✓ 1675ms | ✓ 1367ms | http |
| 45.136.131.54:8448 | ✓ 716ms | ✓ 901ms | ✓ 694ms | ✓ 1065ms | ✓ 1345ms | http |
| 167.103.34.108:8800 | ✓ 1350ms | 否 | ✓ 1140ms | ✓ 1339ms | ✓ 1275ms | http |
| 104.168.158.236:10808 | 否 | 否 | ✓ 1307ms | ✓ 1929ms | ✓ 1807ms | http |
| 137.220.151.110:6005 | ✓ 769ms | 否 | ✓ 866ms | ✓ 1107ms | ✓ 887ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1442ms | ✓ 996ms | 否 | ✓ 988ms | http |
| 137.220.150.22:6005 | ✓ 1656ms | 否 | ✓ 1091ms | 否 | ✓ 1103ms | http |
| 106.75.15.167:7890 | ✓ 1834ms | ✓ 1665ms | 否 | ✓ 1389ms | 否 | http |
| 35.225.22.61:80 | ✓ 989ms | ✓ 1158ms | ✓ 988ms | ✓ 1320ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1979ms | ✓ 1377ms | 否 | 否 | ✓ 1339ms | http |
| 167.103.31.122:8800 | ✓ 1347ms | 否 | ✓ 1448ms | ✓ 1635ms | ✓ 1631ms | http |
| 101.43.127.100:8877 | ✓ 978ms | ✓ 1169ms | ✓ 949ms | ✓ 1236ms | ✓ 931ms | http |
| 1.231.81.166:3128 | ✓ 1628ms | ✓ 916ms | ✓ 1340ms | ✓ 1234ms | ✓ 1135ms | http |
| 219.117.204.211:7799 | ✓ 1602ms | ✓ 1562ms | ✓ 1362ms | 否 | 否 | http |
| 91.238.105.64:2024 | ✓ 1073ms | 否 | ✓ 1965ms | 否 | ✓ 1897ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1964ms | ✓ 1009ms | ✓ 1259ms | ✓ 1002ms | http |
| 38.34.178.241:8444 | 否 | ✓ 874ms | ✓ 176ms | ✓ 794ms | ✓ 591ms | http |
| 194.67.99.223:1080 | ✓ 964ms | 否 | ✓ 1895ms | 否 | ✓ 1454ms | http |
| 2.56.173.45:10808 | ✓ 1602ms | 否 | ✓ 1810ms | 否 | ✓ 1779ms | http |
| 24.144.86.173:1080 | ✓ 686ms | 否 | ✓ 1754ms | ✓ 1294ms | ✓ 1526ms | http |
| 101.47.73.135:3128 | ✓ 1017ms | 否 | ✓ 1712ms | ✓ 1203ms | 否 | http |
| 8.219.97.248:80 | ✓ 1477ms | 否 | ✓ 1575ms | ✓ 1459ms | 否 | http |
| 43.165.195.107:3128 | ✓ 1561ms | 否 | ✓ 937ms | ✓ 1192ms | 否 | http |
| 38.34.179.101:8446 | ✓ 1691ms | 否 | ✓ 145ms | ✓ 874ms | ✓ 1625ms | http |
| 103.156.17.35:8181 | ✓ 1551ms | 否 | 否 | ✓ 1487ms | ✓ 1439ms | http |
| 114.237.77.244:1080 | 否 | ✓ 1657ms | ✓ 921ms | ✓ 1312ms | ✓ 977ms | http |
| 88.80.150.82:8080 | ✓ 1240ms | 否 | ✓ 1167ms | 否 | ✓ 1699ms | https |
| 115.231.181.40:8128 | ✓ 1642ms | ✓ 1787ms | 否 | ✓ 1605ms | 否 | http |
| 45.136.130.180:8445 | ✓ 339ms | ✓ 703ms | ✓ 1356ms | 否 | ✓ 578ms | http |
| 38.34.179.104:8446 | ✓ 1824ms | ✓ 961ms | ✓ 501ms | ✓ 1182ms | ✓ 1788ms | http |
| 38.34.179.100:8450 | ✓ 494ms | ✓ 1989ms | ✓ 150ms | ✓ 779ms | ✓ 912ms | http |
| 45.136.130.179:8445 | ✓ 797ms | ✓ 1690ms | 否 | ✓ 1147ms | ✓ 806ms | http |
| 137.220.150.104:6005 | ✓ 795ms | 否 | ✓ 1169ms | ✓ 1505ms | ✓ 1682ms | http |
| 101.132.61.121:8888 | ✓ 1304ms | ✓ 1250ms | ✓ 1307ms | ✓ 1441ms | ✓ 1258ms | http |
| 142.171.224.229:7890 | ✓ 871ms | ✓ 1388ms | ✓ 780ms | ✓ 833ms | ✓ 555ms | http |
| 38.34.179.98:8453 | ✓ 312ms | ✓ 1859ms | ✓ 407ms | ✓ 739ms | ✓ 734ms | http |
| 38.34.179.86:8452 | ✓ 1546ms | 否 | ✓ 579ms | 否 | ✓ 1665ms | http |
| 166.88.55.83:7890 | ✓ 653ms | ✓ 1103ms | ✓ 651ms | ✓ 827ms | ✓ 656ms | http |
| 122.233.92.199:1111 | ✓ 927ms | ✓ 1170ms | ✓ 1937ms | ✓ 1384ms | ✓ 1654ms | http |
| 202.141.161.53:30001 | ✓ 1249ms | ✓ 1426ms | 否 | ✓ 1453ms | 否 | http |
| 223.16.170.103:80 | ✓ 908ms | 否 | ✓ 863ms | 否 | ✓ 1109ms | http |
| 45.186.6.104:3128 | ✓ 1613ms | ✓ 1978ms | ✓ 1878ms | 否 | 否 | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 650ms | ✓ 1813ms | ✓ 665ms | http |
| 103.39.51.190:8080 | ✓ 1782ms | 否 | 否 | ✓ 1896ms | ✓ 1655ms | http |
| 94.16.114.3:40000 | ✓ 1136ms | 否 | ✓ 1367ms | 否 | ✓ 1845ms | http |
| 144.31.79.117:8888 | ✓ 1185ms | 否 | ✓ 1872ms | 否 | ✓ 1756ms | http |
| 45.136.130.178:8449 | ✓ 1269ms | 否 | ✓ 840ms | 否 | ✓ 814ms | http |
| 38.34.178.186:8451 | ✓ 273ms | 否 | ✓ 918ms | ✓ 1381ms | 否 | http |
| 34.96.238.40:8080 | 否 | ✓ 1148ms | 否 | ✓ 1038ms | ✓ 1063ms | http |
| 38.34.179.39:8452 | 否 | ✓ 786ms | ✓ 560ms | ✓ 1884ms | ✓ 935ms | http |
| 38.34.178.7:8452 | 否 | 否 | ✓ 202ms | ✓ 754ms | ✓ 617ms | http |
| 38.34.179.57:8453 | 否 | 否 | ✓ 183ms | ✓ 792ms | ✓ 692ms | http |
| 180.125.216.109:8118 | ✓ 1619ms | ✓ 1884ms | 否 | 否 | ✓ 1956ms | http |
| 38.34.179.51:8445 | 否 | 否 | ✓ 480ms | ✓ 747ms | ✓ 1505ms | http |
| 207.254.71.62:8088 | ✓ 1100ms | 否 | 否 | ✓ 1882ms | ✓ 1846ms | http |
| 103.82.23.118:5216 | ✓ 1499ms | 否 | ✓ 1187ms | 否 | ✓ 1548ms | http |
| 210.223.44.230:3128 | ✓ 625ms | ✓ 1090ms | ✓ 1157ms | ✓ 949ms | ✓ 708ms | http |
| 38.34.179.150:8449 | ✓ 1201ms | 否 | ✓ 358ms | ✓ 762ms | ✓ 882ms | http |
| 72.56.79.129:1080 | ✓ 941ms | ✓ 1800ms | ✓ 1422ms | ✓ 1967ms | ✓ 1562ms | http |
| 103.69.84.106:3131 | 否 | 否 | ✓ 1230ms | ✓ 1158ms | ✓ 946ms | http |
| 137.184.6.37:3128 | ✓ 397ms | ✓ 875ms | ✓ 793ms | ✓ 744ms | ✓ 578ms | http |
| 157.0.142.246:10057 | 否 | 否 | ✓ 1369ms | ✓ 1728ms | ✓ 1079ms | http |
| 157.230.38.173:3128 | ✓ 1504ms | 否 | ✓ 734ms | ✓ 1388ms | ✓ 1390ms | http |
| 158.101.113.18:80 | ✓ 516ms | 否 | ✓ 1946ms | ✓ 1589ms | ✓ 1523ms | http |
| 45.136.198.40:3128 | ✓ 1250ms | ✓ 1933ms | ✓ 945ms | 否 | 否 | http |
| 106.117.208.101:7890 | ✓ 1046ms | ✓ 1349ms | ✓ 1099ms | 否 | ✓ 1059ms | http |
| 59.8.203.55:80 | ✓ 1182ms | 否 | 否 | ✓ 1319ms | ✓ 808ms | http |
| 121.230.9.252:1080 | ✓ 1413ms | ✓ 1999ms | ✓ 1957ms | 否 | 否 | http |
| 116.80.96.103:3172 | ✓ 489ms | ✓ 1989ms | ✓ 487ms | ✓ 819ms | ✓ 643ms | http |
| 43.99.54.236:5555 | ✓ 679ms | ✓ 1030ms | ✓ 675ms | ✓ 854ms | ✓ 679ms | http |

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
