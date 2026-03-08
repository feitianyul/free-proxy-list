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

最后更新：2026-03-08 21:21:27 UTC（2026-03-09 05:21:27 UTC+8）

**代理总数：65**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 65 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 536ms | ✓ 1916ms | ✓ 1648ms | ✓ 1348ms | ✓ 1024ms | http |
| 103.113.70.189:1081 | 否 | ✓ 944ms | 否 | ✓ 1169ms | ✓ 892ms | http |
| 61.72.110.54:3128 | ✓ 823ms | ✓ 1553ms | ✓ 1153ms | ✓ 1276ms | ✓ 926ms | http |
| 211.171.114.154:3128 | ✓ 976ms | ✓ 1322ms | ✓ 1927ms | ✓ 1526ms | ✓ 1395ms | http |
| 165.227.5.10:8888 | 否 | ✓ 1227ms | ✓ 1399ms | 否 | ✓ 1777ms | http |
| 61.72.221.194:3128 | ✓ 823ms | ✓ 1335ms | ✓ 1317ms | ✓ 1152ms | ✓ 1660ms | http |
| 62.113.119.14:8080 | ✓ 1119ms | ✓ 1959ms | 否 | ✓ 1453ms | ✓ 1176ms | http |
| 121.128.121.54:3128 | ✓ 1785ms | ✓ 1436ms | ✓ 906ms | ✓ 1117ms | 否 | http |
| 194.213.18.200:443 | ✓ 993ms | ✓ 1906ms | 否 | 否 | ✓ 1975ms | http |
| 1.231.81.166:3128 | ✓ 1794ms | ✓ 1542ms | ✓ 1358ms | ✓ 1423ms | ✓ 837ms | http |
| 14.56.107.244:3128 | ✓ 1869ms | 否 | ✓ 1290ms | ✓ 1399ms | ✓ 883ms | http |
| 152.42.213.210:8080 | ✓ 1661ms | 否 | ✓ 1553ms | ✓ 1618ms | ✓ 1954ms | http |
| 35.225.22.61:80 | ✓ 246ms | ✓ 1215ms | 否 | ✓ 1012ms | ✓ 743ms | http |
| 202.155.12.161:443 | ✓ 1684ms | ✓ 1464ms | ✓ 1204ms | ✓ 1259ms | ✓ 1062ms | http |
| 61.72.221.94:3128 | ✓ 1286ms | ✓ 1335ms | ✓ 717ms | ✓ 1154ms | 否 | http |
| 81.70.169.194:80 | ✓ 1076ms | ✓ 1488ms | ✓ 1167ms | ✓ 1484ms | ✓ 1148ms | http |
| 101.43.255.96:80 | ✓ 1132ms | ✓ 1494ms | ✓ 1103ms | ✓ 1615ms | ✓ 1075ms | http |
| 178.236.245.17:3128 | ✓ 1090ms | 否 | ✓ 964ms | ✓ 1602ms | ✓ 1259ms | http |
| 5.129.206.247:8888 | ✓ 1084ms | 否 | ✓ 1449ms | 否 | ✓ 1997ms | http |
| 185.115.74.185:8080 | ✓ 1067ms | ✓ 1930ms | ✓ 1508ms | 否 | 否 | http |
| 115.231.181.40:8128 | ✓ 1242ms | 否 | ✓ 1012ms | ✓ 1612ms | ✓ 1824ms | http |
| 190.9.109.198:999 | ✓ 814ms | ✓ 1453ms | ✓ 1127ms | ✓ 1349ms | ✓ 1129ms | http |
| 190.9.109.207:999 | ✓ 713ms | ✓ 1445ms | ✓ 1188ms | ✓ 1497ms | 否 | http |
| 47.95.231.180:8084 | ✓ 786ms | ✓ 1119ms | ✓ 796ms | ✓ 1094ms | ✓ 902ms | http |
| 103.149.99.128:3128 | ✓ 1701ms | ✓ 1728ms | ✓ 1477ms | ✓ 1373ms | ✓ 1358ms | http |
| 103.82.23.118:5196 | ✓ 1710ms | 否 | ✓ 1354ms | ✓ 1871ms | ✓ 1345ms | http |
| 47.105.98.23:3128 | 否 | ✓ 1555ms | ✓ 899ms | 否 | ✓ 975ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1515ms | ✓ 1659ms | ✓ 1814ms | http |
| 103.82.23.118:5249 | ✓ 1716ms | 否 | ✓ 1834ms | 否 | ✓ 1486ms | http |
| 175.215.54.252:3040 | 否 | ✓ 1609ms | ✓ 1425ms | ✓ 1406ms | ✓ 1432ms | http |
| 168.235.110.63:3128 | ✓ 260ms | ✓ 1873ms | ✓ 1015ms | ✓ 1055ms | ✓ 764ms | http |
| 103.35.188.243:3128 | 否 | ✓ 1884ms | 否 | ✓ 1164ms | ✓ 841ms | http |
| 8.222.175.80:6128 | ✓ 1012ms | ✓ 1993ms | ✓ 959ms | ✓ 1980ms | ✓ 1118ms | http |
| 45.140.147.155:1081 | ✓ 981ms | ✓ 1445ms | ✓ 1184ms | ✓ 1526ms | ✓ 1147ms | http |
| 14.103.139.54:10808 | ✓ 954ms | ✓ 1105ms | ✓ 990ms | ✓ 1279ms | ✓ 970ms | http |
| 170.78.208.245:999 | ✓ 1774ms | ✓ 1577ms | ✓ 1758ms | ✓ 1494ms | ✓ 1297ms | http |
| 185.243.218.43:49153 | ✓ 993ms | 否 | ✓ 1777ms | ✓ 1973ms | ✓ 1728ms | http |
| 47.110.42.192:9003 | 否 | 否 | ✓ 1678ms | ✓ 1801ms | ✓ 1596ms | http |
| 103.215.36.88:17621 | 否 | ✓ 1461ms | ✓ 1086ms | ✓ 1411ms | ✓ 1204ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 1514ms | ✓ 1736ms | ✓ 1527ms | http |
| 88.80.150.82:8080 | ✓ 1734ms | 否 | 否 | ✓ 1983ms | ✓ 1919ms | https |
| 120.238.159.234:22222 | ✓ 1022ms | ✓ 1376ms | ✓ 1246ms | ✓ 1257ms | ✓ 988ms | http |
| 178.236.245.59:3128 | ✓ 687ms | 否 | ✓ 1098ms | 否 | ✓ 1257ms | http |
| 27.73.57.47:10018 | 否 | 否 | ✓ 1697ms | ✓ 1838ms | ✓ 1890ms | http |
| 103.215.36.88:16894 | ✓ 1109ms | ✓ 1474ms | ✓ 1134ms | ✓ 1338ms | ✓ 1059ms | http |
| 67.169.98.211:443 | ✓ 1323ms | ✓ 1848ms | 否 | ✓ 1870ms | 否 | http |
| 147.45.251.242:8888 | ✓ 1165ms | 否 | ✓ 1832ms | 否 | ✓ 1908ms | http |
| 121.177.104.201:3192 | ✓ 1951ms | 否 | ✓ 1613ms | ✓ 1850ms | 否 | http |
| 46.249.103.192:443 | ✓ 1602ms | 否 | 否 | ✓ 1967ms | ✓ 1750ms | http |
| 180.76.115.231:3128 | ✓ 1108ms | ✓ 1420ms | ✓ 1025ms | ✓ 1476ms | ✓ 1285ms | http |
| 103.82.23.118:5171 | ✓ 1443ms | 否 | ✓ 1328ms | ✓ 1586ms | ✓ 1789ms | http |
| 91.233.223.147:3128 | ✓ 859ms | 否 | ✓ 814ms | ✓ 1896ms | ✓ 1514ms | http |
| 120.92.212.16:7890 | ✓ 1043ms | 否 | 否 | ✓ 1366ms | ✓ 1849ms | http |
| 120.92.212.16:8890 | ✓ 1093ms | ✓ 1313ms | ✓ 1062ms | 否 | 否 | http |
| 116.80.82.232:3172 | ✓ 1810ms | 否 | ✓ 1708ms | ✓ 1951ms | ✓ 1775ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1388ms | ✓ 1436ms | ✓ 1504ms | http |
| 116.80.82.218:3172 | ✓ 1650ms | 否 | ✓ 1658ms | 否 | ✓ 1779ms | http |
| 45.136.198.40:3128 | ✓ 1334ms | ✓ 1702ms | ✓ 1905ms | 否 | ✓ 1950ms | http |
| 152.42.213.210:80 | 否 | 否 | ✓ 1667ms | ✓ 1162ms | ✓ 1107ms | http |
| 103.215.36.88:18476 | ✓ 1116ms | ✓ 1281ms | ✓ 1139ms | ✓ 1458ms | ✓ 1170ms | http |
| 172.212.68.37:3128 | ✓ 1477ms | ✓ 1617ms | ✓ 761ms | ✓ 1665ms | ✓ 1782ms | http |
| 45.186.6.104:3128 | ✓ 1323ms | ✓ 1843ms | ✓ 1628ms | 否 | 否 | http |
| 51.79.207.21:8080 | 否 | 否 | ✓ 1825ms | ✓ 1889ms | ✓ 1253ms | http |
| 162.248.165.72:1080 | ✓ 1344ms | 否 | ✓ 1810ms | ✓ 1364ms | ✓ 1925ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1530ms | ✓ 1854ms | ✓ 1715ms | ✓ 1129ms | http |

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
