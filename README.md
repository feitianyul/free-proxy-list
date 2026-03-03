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

最后更新：2026-03-03 03:30:33 UTC（2026-03-03 11:30:33 UTC+8）

**代理总数：76**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 3.213.157.4:3128 | ✓ 234ms | 否 | ✓ 1177ms | ✓ 1371ms | ✓ 1278ms | http |
| 166.0.192.117:8888 | ✓ 1986ms | 否 | ✓ 363ms | 否 | ✓ 865ms | http |
| 185.115.74.185:8080 | ✓ 1528ms | ✓ 1932ms | ✓ 1628ms | 否 | 否 | http |
| 205.209.118.30:3138 | ✓ 749ms | ✓ 984ms | ✓ 1919ms | ✓ 1142ms | ✓ 915ms | http |
| 83.219.250.8:62920 | ✓ 596ms | ✓ 1972ms | ✓ 1579ms | 否 | ✓ 1576ms | http |
| 186.148.180.46:999 | ✓ 1244ms | ✓ 1850ms | ✓ 1558ms | 否 | ✓ 1345ms | http |
| 115.231.181.40:8128 | ✓ 1033ms | ✓ 1157ms | ✓ 1847ms | 否 | 否 | http |
| 20.78.118.91:8561 | ✓ 1591ms | ✓ 1622ms | ✓ 993ms | ✓ 1189ms | 否 | http |
| 142.171.85.32:1080 | ✓ 1452ms | ✓ 1226ms | ✓ 1571ms | ✓ 1225ms | 否 | http |
| 20.210.39.153:8561 | ✓ 1606ms | ✓ 1624ms | ✓ 992ms | ✓ 1195ms | 否 | http |
| 20.27.14.220:8561 | ✓ 1778ms | ✓ 1963ms | ✓ 1499ms | 否 | ✓ 1622ms | http |
| 95.85.252.153:21064 | ✓ 992ms | 否 | ✓ 1109ms | ✓ 1899ms | 否 | http |
| 20.210.76.104:8561 | ✓ 1288ms | 否 | ✓ 878ms | ✓ 1198ms | ✓ 1212ms | http |
| 14.56.107.244:3128 | ✓ 1705ms | 否 | ✓ 1732ms | ✓ 1602ms | 否 | http |
| 20.210.76.178:8561 | ✓ 597ms | ✓ 1035ms | ✓ 667ms | ✓ 946ms | ✓ 769ms | http |
| 81.70.169.194:80 | ✓ 1157ms | ✓ 1408ms | ✓ 1019ms | ✓ 1471ms | ✓ 1103ms | http |
| 101.43.255.96:80 | ✓ 1052ms | 否 | ✓ 1054ms | 否 | ✓ 1076ms | http |
| 120.92.212.16:7890 | ✓ 1006ms | ✓ 1266ms | ✓ 1552ms | 否 | 否 | http |
| 115.76.5.32:10008 | ✓ 1801ms | 否 | 否 | ✓ 1929ms | ✓ 1737ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1266ms | ✓ 1967ms | ✓ 1695ms | http |
| 120.92.212.16:8890 | ✓ 1512ms | 否 | ✓ 1242ms | 否 | ✓ 1824ms | http |
| 35.234.17.221:8080 | 否 | ✓ 1505ms | ✓ 997ms | 否 | ✓ 1213ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 969ms | ✓ 982ms | ✓ 820ms | http |
| 20.27.15.111:8561 | ✓ 675ms | ✓ 1271ms | ✓ 572ms | ✓ 879ms | ✓ 696ms | http |
| 45.125.67.37:8443 | ✓ 1033ms | 否 | ✓ 1027ms | 否 | ✓ 1092ms | http |
| 20.27.11.248:8561 | ✓ 675ms | ✓ 1817ms | ✓ 547ms | ✓ 885ms | ✓ 740ms | http |
| 121.128.121.54:3128 | ✓ 1786ms | ✓ 1458ms | ✓ 700ms | ✓ 1111ms | ✓ 820ms | http |
| 91.238.104.171:2023 | ✓ 1860ms | 否 | 否 | ✓ 1679ms | ✓ 1294ms | http |
| 14.56.177.44:3128 | ✓ 696ms | ✓ 1240ms | ✓ 1911ms | ✓ 1657ms | 否 | http |
| 150.107.140.238:3128 | ✓ 1976ms | 否 | 否 | ✓ 1267ms | ✓ 1518ms | http |
| 61.72.221.94:3128 | ✓ 1676ms | ✓ 1790ms | ✓ 1429ms | 否 | 否 | http |
| 115.76.5.32:10007 | ✓ 1456ms | 否 | 否 | ✓ 1720ms | ✓ 1700ms | http |
| 177.243.209.133:999 | ✓ 1267ms | 否 | ✓ 773ms | ✓ 1445ms | ✓ 1160ms | http |
| 160.238.65.5:3128 | 否 | ✓ 1891ms | 否 | ✓ 1333ms | ✓ 1419ms | http |
| 20.78.26.206:8561 | ✓ 1287ms | ✓ 1458ms | ✓ 1181ms | ✓ 1461ms | ✓ 1265ms | http |
| 91.99.99.83:9000 | 否 | ✓ 1733ms | ✓ 1611ms | 否 | ✓ 1945ms | http |
| 20.27.15.49:8561 | ✓ 1720ms | ✓ 1854ms | ✓ 1094ms | ✓ 1553ms | ✓ 1441ms | http |
| 20.210.76.175:8561 | ✓ 1720ms | ✓ 1880ms | ✓ 1083ms | ✓ 1613ms | ✓ 1488ms | http |
| 5.75.196.26:40000 | ✓ 518ms | 否 | ✓ 791ms | 否 | ✓ 1382ms | http |
| 125.128.12.144:3128 | ✓ 681ms | 否 | 否 | ✓ 1080ms | ✓ 1938ms | http |
| 162.240.154.26:3128 | ✓ 1004ms | ✓ 1258ms | ✓ 1542ms | ✓ 1296ms | ✓ 1045ms | http |
| 154.90.48.209:9090 | ✓ 1887ms | 否 | 否 | ✓ 1337ms | ✓ 1740ms | http |
| 160.238.65.4:3128 | ✓ 1575ms | 否 | 否 | ✓ 1370ms | ✓ 1065ms | http |
| 69.164.194.19:3128 | ✓ 562ms | ✓ 1586ms | ✓ 192ms | ✓ 1564ms | ✓ 1172ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 832ms | ✓ 1041ms | ✓ 863ms | http |
| 121.230.8.109:1080 | ✓ 1301ms | ✓ 1376ms | ✓ 1237ms | ✓ 1833ms | ✓ 1413ms | http |
| 43.134.166.79:8888 | ✓ 1658ms | 否 | ✓ 1244ms | ✓ 1495ms | ✓ 1387ms | http |
| 45.88.0.98:3128 | 否 | 否 | ✓ 1238ms | ✓ 1371ms | ✓ 1411ms | http |
| 125.128.12.14:3128 | ✓ 1265ms | ✓ 1456ms | ✓ 1241ms | ✓ 1086ms | ✓ 900ms | http |
| 61.72.221.234:3128 | 否 | ✓ 1586ms | ✓ 1288ms | 否 | ✓ 897ms | http |
| 62.113.119.14:8080 | ✓ 729ms | 否 | ✓ 747ms | ✓ 1517ms | ✓ 1144ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 1462ms | ✓ 1441ms | ✓ 801ms | http |
| 46.249.103.192:443 | ✓ 1187ms | 否 | ✓ 1741ms | 否 | ✓ 1810ms | http |
| 188.166.208.168:9876 | 否 | 否 | ✓ 1188ms | ✓ 1119ms | ✓ 916ms | http |
| 103.215.36.88:19328 | ✓ 1062ms | ✓ 1463ms | ✓ 1108ms | ✓ 1466ms | ✓ 1040ms | http |
| 43.165.195.107:3128 | ✓ 1584ms | 否 | ✓ 1211ms | ✓ 1307ms | ✓ 1373ms | http |
| 202.129.206.239:3128 | 否 | 否 | ✓ 1838ms | ✓ 1692ms | ✓ 1417ms | http |
| 157.0.142.246:10061 | ✓ 1059ms | 否 | 否 | ✓ 1427ms | ✓ 1103ms | http |
| 160.238.65.9:3128 | ✓ 1884ms | 否 | ✓ 1514ms | ✓ 1743ms | 否 | http |
| 45.136.198.40:3128 | ✓ 1158ms | 否 | ✓ 1640ms | ✓ 1929ms | ✓ 1801ms | http |
| 45.129.141.143:3128 | ✓ 931ms | 否 | ✓ 1884ms | 否 | ✓ 1743ms | http |
| 47.77.180.205:1080 | ✓ 592ms | ✓ 1284ms | ✓ 523ms | ✓ 860ms | ✓ 648ms | http |
| 120.79.99.232:8099 | ✓ 1190ms | ✓ 1449ms | ✓ 1238ms | ✓ 1353ms | ✓ 1133ms | http |
| 103.215.36.88:19626 | ✓ 1121ms | 否 | ✓ 1711ms | ✓ 1541ms | ✓ 1158ms | http |
| 160.250.5.22:1 | ✓ 1797ms | 否 | ✓ 1348ms | ✓ 1286ms | ✓ 1014ms | http |
| 103.82.23.118:5203 | ✓ 1593ms | 否 | ✓ 1106ms | ✓ 1879ms | ✓ 1435ms | http |
| 160.250.4.245:1 | ✓ 1784ms | 否 | ✓ 1564ms | ✓ 1420ms | ✓ 1039ms | http |
| 195.123.209.48:3128 | ✓ 1384ms | 否 | ✓ 1910ms | ✓ 1957ms | ✓ 1806ms | http |
| 74.208.234.198:443 | ✓ 1844ms | ✓ 1262ms | 否 | ✓ 1679ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1786ms | 否 | 否 | ✓ 1489ms | ✓ 1462ms | http |
| 2.56.178.131:443 | ✓ 1337ms | 否 | ✓ 1608ms | 否 | ✓ 1731ms | http |
| 121.230.8.36:1080 | ✓ 1384ms | 否 | ✓ 1249ms | ✓ 1586ms | 否 | http |
| 103.215.36.88:15247 | ✓ 1037ms | ✓ 1442ms | ✓ 1229ms | ✓ 1391ms | 否 | http |
| 186.116.148.52:8080 | ✓ 1430ms | 否 | 否 | ✓ 1904ms | ✓ 1728ms | http |
| 111.79.111.126:3128 | ✓ 1316ms | 否 | ✓ 1972ms | 否 | ✓ 1536ms | http |
| 45.140.147.82:1082 | ✓ 596ms | ✓ 1154ms | ✓ 786ms | ✓ 1599ms | 否 | http |

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
