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

最后更新：2026-04-12 06:22:49 UTC（2026-04-12 14:22:49 UTC+8）

**代理总数：62**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 892ms | 否 | ✓ 1031ms | ✓ 1027ms | ✓ 872ms | http |
| 147.161.239.240:8800 | ✓ 1183ms | ✓ 1841ms | ✓ 1494ms | ✓ 1822ms | ✓ 1388ms | http |
| 113.160.132.26:8080 | ✓ 1660ms | ✓ 1781ms | 否 | 否 | ✓ 949ms | http |
| 167.103.115.102:8800 | ✓ 1804ms | 否 | ✓ 975ms | 否 | ✓ 1204ms | http |
| 202.141.161.53:10808 | ✓ 1065ms | ✓ 1336ms | ✓ 1226ms | ✓ 1215ms | ✓ 1098ms | http |
| 5.196.101.18:3128 | ✓ 1932ms | ✓ 1839ms | ✓ 1257ms | 否 | 否 | http |
| 167.103.144.127:8800 | ✓ 1175ms | 否 | ✓ 1394ms | 否 | ✓ 1479ms | http |
| 167.103.34.108:8800 | ✓ 1547ms | 否 | ✓ 1488ms | ✓ 1673ms | 否 | http |
| 167.103.31.122:8800 | ✓ 1689ms | 否 | ✓ 1523ms | ✓ 1934ms | 否 | http |
| 74.50.96.247:8888 | ✓ 482ms | ✓ 1003ms | ✓ 988ms | 否 | 否 | http |
| 45.167.125.21:999 | ✓ 1021ms | 否 | ✓ 1296ms | ✓ 1835ms | ✓ 1534ms | http |
| 217.76.245.80:999 | 否 | 否 | ✓ 1402ms | ✓ 1657ms | ✓ 1202ms | http |
| 162.240.154.26:3128 | ✓ 1791ms | ✓ 1994ms | ✓ 1609ms | ✓ 1838ms | 否 | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 1303ms | ✓ 1252ms | ✓ 1068ms | http |
| 159.223.225.118:8888 | ✓ 1190ms | 否 | ✓ 1946ms | ✓ 1771ms | ✓ 1560ms | http |
| 103.157.200.126:3128 | 否 | 否 | ✓ 1424ms | ✓ 1861ms | ✓ 1482ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1578ms | ✓ 1367ms | ✓ 1855ms | ✓ 1435ms | http |
| 5.104.87.17:8051 | ✓ 1012ms | 否 | ✓ 1404ms | ✓ 1593ms | ✓ 1090ms | http |
| 185.76.241.110:10001 | ✓ 1557ms | 否 | ✓ 770ms | 否 | ✓ 1606ms | http |
| 185.76.240.167:10001 | ✓ 1553ms | 否 | ✓ 769ms | 否 | ✓ 1622ms | http |
| 160.119.69.7:8080 | ✓ 1524ms | 否 | ✓ 1294ms | 否 | ✓ 1812ms | http |
| 185.76.241.94:10001 | ✓ 1521ms | 否 | ✓ 1165ms | 否 | ✓ 1618ms | http |
| 185.76.240.244:10001 | ✓ 1522ms | 否 | ✓ 1161ms | 否 | ✓ 1778ms | http |
| 185.76.240.146:10001 | ✓ 1522ms | 否 | ✓ 1161ms | 否 | ✓ 1786ms | http |
| 185.76.240.245:10001 | ✓ 1521ms | 否 | ✓ 1157ms | 否 | ✓ 1909ms | http |
| 185.76.240.143:10001 | ✓ 1518ms | 否 | ✓ 1161ms | 否 | ✓ 1910ms | http |
| 185.76.240.145:10001 | ✓ 1519ms | 否 | ✓ 1652ms | 否 | ✓ 1700ms | http |
| 137.59.47.73:3128 | ✓ 1362ms | ✓ 1389ms | ✓ 1376ms | 否 | ✓ 975ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1604ms | ✓ 1259ms | ✓ 1131ms | http |
| 101.32.244.83:8080 | ✓ 988ms | 否 | ✓ 934ms | ✓ 1258ms | ✓ 1217ms | http |
| 121.43.196.213:8222 | ✓ 937ms | ✓ 1023ms | ✓ 901ms | ✓ 1138ms | ✓ 917ms | http |
| 121.43.196.210:8222 | ✓ 934ms | ✓ 1014ms | ✓ 919ms | ✓ 1149ms | ✓ 924ms | http |
| 114.55.226.123:10086 | ✓ 1101ms | ✓ 1466ms | ✓ 1155ms | ✓ 1271ms | ✓ 1066ms | http |
| 103.82.23.118:5195 | ✓ 1761ms | 否 | ✓ 1597ms | 否 | ✓ 1530ms | http |
| 8.219.97.248:80 | ✓ 971ms | ✓ 1855ms | 否 | ✓ 1453ms | 否 | http |
| 59.46.216.131:30001 | ✓ 993ms | 否 | ✓ 1116ms | 否 | ✓ 1028ms | http |
| 185.76.240.254:10001 | ✓ 762ms | 否 | ✓ 1103ms | 否 | ✓ 1873ms | http |
| 115.231.181.40:8128 | ✓ 1233ms | 否 | ✓ 1207ms | 否 | ✓ 1923ms | http |
| 46.30.46.133:3128 | ✓ 1756ms | 否 | ✓ 1609ms | ✓ 1838ms | ✓ 1407ms | http |
| 80.250.165.242:3128 | ✓ 1143ms | 否 | ✓ 1464ms | 否 | ✓ 1830ms | http |
| 185.76.240.238:10001 | ✓ 1182ms | 否 | ✓ 1370ms | 否 | ✓ 1871ms | http |
| 185.76.240.139:10001 | ✓ 997ms | 否 | ✓ 1498ms | 否 | ✓ 1888ms | http |
| 83.219.250.8:62920 | ✓ 937ms | 否 | ✓ 1035ms | 否 | ✓ 1445ms | http |
| 103.113.70.189:1081 | 否 | 否 | ✓ 1599ms | ✓ 1567ms | ✓ 883ms | http |
| 185.76.240.142:10001 | ✓ 1442ms | 否 | ✓ 1223ms | 否 | ✓ 1976ms | http |
| 150.241.116.228:3128 | ✓ 1664ms | 否 | ✓ 1647ms | 否 | ✓ 1655ms | http |
| 185.76.240.169:10001 | 否 | 否 | ✓ 1500ms | ✓ 1987ms | ✓ 1973ms | http |
| 107.172.102.234:40621 | ✓ 523ms | 否 | ✓ 1346ms | 否 | ✓ 1932ms | http |
| 34.101.184.164:3128 | ✓ 1446ms | 否 | ✓ 802ms | ✓ 1382ms | ✓ 1394ms | http |
| 212.58.132.5:8888 | 否 | 否 | ✓ 1518ms | ✓ 1507ms | ✓ 1480ms | http |
| 94.72.109.214:8888 | ✓ 1959ms | 否 | ✓ 1607ms | ✓ 1655ms | ✓ 1231ms | http |
| 152.32.132.190:7890 | ✓ 896ms | 否 | 否 | ✓ 1282ms | ✓ 1403ms | http |
| 122.144.6.42:3888 | ✓ 1743ms | 否 | ✓ 1336ms | 否 | ✓ 1780ms | http |
| 116.171.106.111:3443 | 否 | ✓ 1535ms | ✓ 1341ms | ✓ 1776ms | ✓ 1293ms | http |
| 101.43.127.100:8877 | ✓ 1004ms | ✓ 1059ms | ✓ 1101ms | 否 | ✓ 1583ms | http |
| 147.45.180.91:8888 | ✓ 1186ms | 否 | ✓ 1935ms | 否 | ✓ 1772ms | http |
| 222.228.171.92:8080 | ✓ 759ms | ✓ 1936ms | ✓ 1321ms | 否 | ✓ 881ms | http |
| 103.67.46.225:3125 | ✓ 1981ms | 否 | 否 | ✓ 1683ms | ✓ 1586ms | http |
| 45.136.130.176:8451 | ✓ 1782ms | ✓ 802ms | ✓ 1453ms | 否 | 否 | http |
| 20.78.213.56:80 | ✓ 682ms | ✓ 1231ms | ✓ 655ms | ✓ 1057ms | ✓ 857ms | http |
| 103.39.51.207:8080 | ✓ 1366ms | 否 | 否 | ✓ 1973ms | ✓ 1889ms | http |
| 38.145.208.184:8443 | ✓ 669ms | ✓ 1597ms | ✓ 212ms | ✓ 966ms | ✓ 581ms | http |

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
