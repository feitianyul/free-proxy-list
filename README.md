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

最后更新：2026-05-03 23:39:32 UTC（2026-05-04 07:39:32 UTC+8）

**代理总数：71**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 71 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 71 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 46.105.190.40:3128 | ✓ 1110ms | ✓ 1249ms | ✓ 573ms | ✓ 1730ms | ✓ 1396ms | http |
| 47.85.51.197:1080 | 否 | ✓ 1087ms | ✓ 1089ms | 否 | ✓ 773ms | http |
| 154.64.232.35:8080 | 否 | ✓ 1945ms | ✓ 1075ms | ✓ 1777ms | ✓ 1007ms | http |
| 206.206.126.177:2412 | ✓ 850ms | ✓ 1519ms | 否 | ✓ 1201ms | ✓ 1197ms | http |
| 45.167.124.71:999 | ✓ 1255ms | ✓ 1817ms | ✓ 1204ms | ✓ 1707ms | ✓ 1419ms | http |
| 45.59.122.132:80 | ✓ 857ms | ✓ 1480ms | ✓ 800ms | ✓ 1750ms | ✓ 1373ms | http |
| 1.231.81.166:3128 | ✓ 901ms | ✓ 1038ms | ✓ 1020ms | ✓ 1236ms | ✓ 890ms | http |
| 86.104.74.110:1082 | ✓ 798ms | ✓ 1326ms | ✓ 874ms | ✓ 1891ms | ✓ 1440ms | http |
| 113.160.132.26:8080 | ✓ 1049ms | 否 | ✓ 1050ms | ✓ 1455ms | ✓ 1135ms | http |
| 80.92.204.47:1081 | 否 | 否 | ✓ 1578ms | ✓ 1961ms | ✓ 1235ms | http |
| 37.187.109.70:10111 | ✓ 907ms | ✓ 1778ms | ✓ 1046ms | 否 | 否 | http |
| 173.212.245.136:8888 | ✓ 969ms | 否 | ✓ 1797ms | 否 | ✓ 1837ms | http |
| 45.153.231.229:8080 | ✓ 1774ms | ✓ 1928ms | ✓ 1575ms | 否 | 否 | http |
| 62.113.119.14:8080 | ✓ 1132ms | ✓ 1391ms | ✓ 976ms | ✓ 1495ms | 否 | http |
| 148.230.4.241:999 | ✓ 1202ms | ✓ 1871ms | ✓ 738ms | 否 | 否 | http |
| 194.150.220.163:1082 | ✓ 1136ms | ✓ 1788ms | ✓ 425ms | ✓ 1422ms | ✓ 1567ms | http |
| 103.157.200.126:3128 | ✓ 1273ms | 否 | ✓ 1004ms | 否 | ✓ 1835ms | http |
| 218.108.131.186:17890 | ✓ 949ms | ✓ 1173ms | ✓ 1197ms | 否 | 否 | http |
| 129.213.162.27:17777 | 否 | ✓ 1359ms | 否 | ✓ 1352ms | ✓ 1730ms | http |
| 47.77.216.82:1080 | ✓ 1737ms | ✓ 1003ms | ✓ 321ms | 否 | 否 | http |
| 86.104.74.110:1081 | ✓ 1061ms | ✓ 1536ms | ✓ 706ms | ✓ 1596ms | ✓ 1284ms | http |
| 130.61.174.200:1080 | ✓ 1114ms | 否 | ✓ 745ms | ✓ 1275ms | ✓ 1056ms | http |
| 43.133.44.89:8888 | ✓ 1435ms | 否 | ✓ 827ms | ✓ 1145ms | 否 | http |
| 147.45.178.211:14658 | ✓ 959ms | ✓ 1679ms | ✓ 1365ms | 否 | ✓ 1683ms | http |
| 8.211.166.184:8081 | ✓ 1473ms | ✓ 1235ms | ✓ 680ms | ✓ 983ms | ✓ 745ms | http |
| 91.233.223.147:3128 | ✓ 1107ms | ✓ 1962ms | ✓ 1505ms | 否 | ✓ 1974ms | http |
| 190.12.150.244:999 | ✓ 1412ms | ✓ 1590ms | ✓ 1130ms | ✓ 1718ms | ✓ 1400ms | http |
| 120.92.212.16:8890 | ✓ 1106ms | 否 | ✓ 1068ms | ✓ 1964ms | 否 | http |
| 109.120.156.122:8090 | ✓ 1256ms | 否 | ✓ 1003ms | 否 | ✓ 1669ms | http |
| 212.58.132.5:8888 | ✓ 1253ms | 否 | ✓ 1793ms | ✓ 1725ms | ✓ 1298ms | http |
| 8.154.21.175:3128 | ✓ 1002ms | ✓ 1187ms | ✓ 1016ms | ✓ 1221ms | ✓ 1006ms | http |
| 121.230.9.225:1080 | ✓ 1243ms | ✓ 1385ms | ✓ 1144ms | ✓ 1528ms | ✓ 1209ms | http |
| 116.171.106.111:3443 | 否 | ✓ 1974ms | ✓ 1876ms | 否 | ✓ 1799ms | http |
| 103.156.248.53:8080 | ✓ 1267ms | 否 | 否 | ✓ 1637ms | ✓ 1531ms | http |
| 121.230.8.55:1080 | ✓ 1066ms | ✓ 1374ms | ✓ 1033ms | ✓ 1408ms | ✓ 1259ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1668ms | ✓ 1029ms | ✓ 1284ms | 否 | http |
| 103.97.198.253:8080 | 否 | 否 | ✓ 1536ms | ✓ 1915ms | ✓ 1582ms | http |
| 45.140.147.155:1082 | ✓ 982ms | ✓ 1390ms | ✓ 1513ms | ✓ 1788ms | ✓ 1369ms | http |
| 152.32.132.190:7890 | 否 | ✓ 1665ms | ✓ 1677ms | 否 | ✓ 1579ms | http |
| 154.90.48.209:9090 | 否 | 否 | ✓ 1770ms | ✓ 1834ms | ✓ 1077ms | http |
| 45.125.67.37:8443 | ✓ 1458ms | 否 | ✓ 1659ms | ✓ 1390ms | ✓ 1114ms | http |
| 86.104.72.220:1081 | ✓ 300ms | ✓ 1010ms | ✓ 872ms | ✓ 1015ms | 否 | http |
| 104.128.138.186:1080 | ✓ 1768ms | ✓ 1976ms | ✓ 1905ms | ✓ 1764ms | ✓ 1678ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1623ms | ✓ 1345ms | ✓ 1094ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1334ms | ✓ 1703ms | ✓ 1443ms | http |
| 193.123.250.39:1080 | ✓ 1459ms | 否 | 否 | ✓ 1360ms | ✓ 1071ms | http |
| 20.164.75.153:8080 | ✓ 1809ms | 否 | ✓ 1216ms | 否 | ✓ 1856ms | http |
| 38.188.247.12:999 | ✓ 1361ms | ✓ 1951ms | ✓ 675ms | 否 | ✓ 1573ms | http |
| 3.101.133.120:80 | 否 | ✓ 1535ms | ✓ 1478ms | ✓ 1180ms | ✓ 1229ms | http |
| 94.131.118.129:1082 | ✓ 831ms | ✓ 1235ms | ✓ 1393ms | ✓ 1709ms | ✓ 1507ms | http |
| 20.205.16.149:3128 | ✓ 938ms | ✓ 1508ms | 否 | ✓ 1255ms | ✓ 951ms | http |
| 94.131.118.129:1081 | ✓ 830ms | ✓ 1260ms | ✓ 1366ms | ✓ 1710ms | ✓ 1553ms | http |
| 92.119.127.208:6005 | ✓ 1050ms | ✓ 1623ms | ✓ 1489ms | ✓ 1852ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1150ms | ✓ 1435ms | ✓ 1169ms | ✓ 1604ms | 否 | http |
| 160.19.19.25:8080 | 否 | 否 | ✓ 1469ms | ✓ 1699ms | ✓ 1982ms | http |
| 86.104.72.219:1081 | ✓ 1294ms | ✓ 1164ms | ✓ 92ms | ✓ 997ms | ✓ 879ms | http |
| 103.35.190.69:1082 | ✓ 192ms | ✓ 1072ms | ✓ 1122ms | ✓ 1150ms | 否 | http |
| 47.74.226.8:5001 | 否 | 否 | ✓ 1224ms | ✓ 1372ms | ✓ 1436ms | http |
| 45.129.141.143:3128 | ✓ 1461ms | ✓ 1797ms | ✓ 1601ms | ✓ 1843ms | ✓ 1722ms | http |
| 167.71.196.178:80 | ✓ 1866ms | 否 | ✓ 833ms | ✓ 1212ms | ✓ 939ms | http |
| 62.60.231.71:56608 | ✓ 1222ms | ✓ 1981ms | 否 | 否 | ✓ 1302ms | http |
| 150.249.255.91:3128 | 否 | ✓ 1800ms | ✓ 688ms | ✓ 1717ms | ✓ 842ms | http |
| 62.133.60.126:24558 | ✓ 984ms | 否 | ✓ 931ms | 否 | ✓ 1558ms | http |
| 94.131.118.39:1081 | ✓ 1112ms | ✓ 1700ms | ✓ 1026ms | ✓ 1677ms | ✓ 1241ms | http |
| 86.104.72.219:1082 | ✓ 282ms | ✓ 920ms | ✓ 1082ms | ✓ 1562ms | ✓ 753ms | http |
| 15.204.238.117:3128 | ✓ 1497ms | ✓ 1289ms | ✓ 111ms | ✓ 1145ms | 否 | http |
| 152.42.177.32:8888 | ✓ 1320ms | 否 | ✓ 1077ms | ✓ 1669ms | ✓ 1469ms | http |
| 121.230.8.250:1080 | ✓ 1133ms | 否 | ✓ 1161ms | ✓ 1461ms | ✓ 1312ms | http |
| 61.52.131.172:8443 | ✓ 1048ms | ✓ 1354ms | ✓ 1062ms | ✓ 1341ms | ✓ 1070ms | http |
| 157.230.220.25:4857 | ✓ 674ms | ✓ 1599ms | 否 | ✓ 1179ms | 否 | http |
| 103.172.70.173:8080 | 否 | 否 | ✓ 1990ms | ✓ 1557ms | ✓ 1520ms | http |

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
