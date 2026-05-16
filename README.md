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

最后更新：2026-05-16 22:45:32 UTC（2026-05-17 06:45:32 UTC+8）

**代理总数：59**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 59 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 59 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 113.160.132.26:8080 | ✓ 1400ms | ✓ 1809ms | ✓ 1015ms | ✓ 1329ms | ✓ 1205ms | http |
| 51.161.50.166:3128 | ✓ 1041ms | 否 | ✓ 1480ms | ✓ 1852ms | ✓ 1470ms | http |
| 170.106.136.181:31002 | ✓ 1034ms | 否 | ✓ 1489ms | 否 | ✓ 1531ms | http |
| 114.214.170.41:27890 | ✓ 1091ms | ✓ 1322ms | ✓ 1381ms | ✓ 1338ms | ✓ 1092ms | http |
| 45.125.67.37:8443 | ✓ 1509ms | 否 | ✓ 1155ms | ✓ 1843ms | ✓ 1307ms | http |
| 212.58.132.5:8888 | ✓ 1093ms | ✓ 1997ms | ✓ 1172ms | ✓ 1568ms | ✓ 1210ms | http |
| 185.200.188.234:10001 | ✓ 1479ms | 否 | ✓ 1899ms | 否 | ✓ 1737ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1282ms | ✓ 1151ms | ✓ 1374ms | 否 | http |
| 120.92.212.16:8890 | ✓ 893ms | 否 | 否 | ✓ 1451ms | ✓ 1000ms | http |
| 1.231.81.166:3128 | ✓ 1821ms | 否 | ✓ 1540ms | ✓ 1914ms | ✓ 1303ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1856ms | ✓ 1990ms | ✓ 994ms | 否 | http |
| 91.242.229.129:8092 | 否 | 否 | ✓ 1771ms | ✓ 1808ms | ✓ 1601ms | http |
| 218.108.131.186:17890 | 否 | ✓ 1726ms | ✓ 981ms | ✓ 1262ms | ✓ 928ms | http |
| 62.113.119.14:8080 | ✓ 871ms | ✓ 1706ms | ✓ 1129ms | 否 | 否 | http |
| 167.99.173.119:3128 | ✓ 501ms | ✓ 963ms | ✓ 1646ms | ✓ 1222ms | 否 | http |
| 8.154.21.175:3128 | ✓ 830ms | ✓ 1030ms | ✓ 870ms | ✓ 1080ms | ✓ 893ms | http |
| 148.230.4.241:999 | ✓ 695ms | ✓ 1494ms | ✓ 670ms | ✓ 1555ms | ✓ 1318ms | http |
| 8.219.97.248:80 | ✓ 1564ms | 否 | ✓ 1433ms | ✓ 1688ms | 否 | http |
| 137.59.47.73:3128 | ✓ 1830ms | 否 | ✓ 1868ms | 否 | ✓ 1175ms | http |
| 128.199.254.13:9090 | ✓ 1986ms | 否 | ✓ 1656ms | ✓ 1875ms | 否 | http |
| 150.107.140.238:3128 | ✓ 1684ms | 否 | 否 | ✓ 1746ms | ✓ 995ms | http |
| 49.156.44.116:8080 | ✓ 1569ms | ✓ 1618ms | ✓ 1440ms | ✓ 1781ms | 否 | http |
| 103.21.220.141:3128 | ✓ 684ms | 否 | ✓ 729ms | ✓ 907ms | ✓ 713ms | http |
| 47.105.98.23:3128 | ✓ 840ms | 否 | ✓ 950ms | 否 | ✓ 1425ms | http |
| 160.238.65.7:3128 | 否 | ✓ 1990ms | 否 | ✓ 1523ms | ✓ 1654ms | http |
| 168.110.52.228:3128 | ✓ 442ms | ✓ 1204ms | ✓ 1288ms | 否 | ✓ 785ms | http |
| 185.71.196.92:1080 | ✓ 1914ms | 否 | ✓ 1264ms | 否 | ✓ 1986ms | http |
| 84.47.150.125:1080 | ✓ 1128ms | 否 | ✓ 1518ms | 否 | ✓ 1882ms | http |
| 34.96.238.40:8080 | ✓ 1131ms | 否 | ✓ 1677ms | 否 | ✓ 968ms | http |
| 42.114.172.179:2045 | ✓ 1506ms | 否 | ✓ 1387ms | ✓ 1657ms | ✓ 1498ms | http |
| 115.231.181.40:8128 | ✓ 878ms | 否 | 否 | ✓ 1406ms | ✓ 1186ms | http |
| 120.92.212.16:7890 | ✓ 1877ms | 否 | ✓ 938ms | ✓ 1859ms | 否 | http |
| 64.188.77.221:3128 | ✓ 1276ms | ✓ 1739ms | ✓ 1158ms | 否 | 否 | http |
| 193.181.35.217:8118 | ✓ 1305ms | 否 | ✓ 944ms | 否 | ✓ 1642ms | http |
| 64.188.77.26:3128 | ✓ 1281ms | ✓ 1837ms | ✓ 1176ms | 否 | 否 | http |
| 101.32.243.189:80 | ✓ 1237ms | ✓ 1448ms | ✓ 1210ms | 否 | ✓ 1233ms | http |
| 159.223.41.216:9090 | ✓ 715ms | 否 | ✓ 1009ms | ✓ 1049ms | ✓ 981ms | http |
| 103.165.138.173:8181 | ✓ 1802ms | 否 | ✓ 1109ms | ✓ 1189ms | ✓ 969ms | http |
| 159.89.31.62:8080 | ✓ 1266ms | 否 | ✓ 1816ms | ✓ 1616ms | ✓ 1482ms | http |
| 3.101.133.120:80 | ✓ 156ms | ✓ 1266ms | ✓ 1439ms | ✓ 747ms | ✓ 560ms | http |
| 138.2.239.213:10010 | ✓ 1328ms | 否 | ✓ 1589ms | ✓ 1098ms | ✓ 802ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1849ms | ✓ 1799ms | ✓ 1605ms | http |
| 157.0.142.246:10057 | ✓ 1195ms | ✓ 1265ms | ✓ 1311ms | ✓ 1479ms | ✓ 1046ms | http |
| 103.147.152.12:1080 | 否 | ✓ 1646ms | ✓ 1528ms | ✓ 1726ms | 否 | http |
| 42.114.172.179:2075 | ✓ 1462ms | 否 | ✓ 1542ms | ✓ 1648ms | ✓ 1465ms | http |
| 128.199.114.189:9090 | ✓ 1490ms | 否 | ✓ 1247ms | ✓ 1125ms | ✓ 968ms | http |
| 86.104.72.219:1081 | ✓ 1255ms | 否 | ✓ 1823ms | ✓ 1937ms | ✓ 1818ms | http |
| 45.88.0.114:3128 | ✓ 813ms | 否 | ✓ 1622ms | 否 | ✓ 1641ms | http |
| 103.147.152.12:1095 | 否 | ✓ 1603ms | ✓ 1082ms | 否 | ✓ 1340ms | http |
| 121.230.9.225:1080 | ✓ 1210ms | ✓ 1305ms | ✓ 897ms | ✓ 1410ms | ✓ 1009ms | http |
| 121.230.9.96:1080 | ✓ 1239ms | ✓ 1326ms | ✓ 1182ms | ✓ 1508ms | ✓ 1098ms | http |
| 204.157.185.2:999 | ✓ 1577ms | 否 | ✓ 1564ms | 否 | ✓ 1954ms | http |
| 61.52.131.172:8443 | ✓ 865ms | ✓ 1246ms | ✓ 1047ms | ✓ 1218ms | ✓ 918ms | http |
| 43.156.90.221:10808 | ✓ 1289ms | ✓ 1825ms | 否 | ✓ 995ms | 否 | http |
| 128.199.113.85:9090 | ✓ 882ms | 否 | ✓ 726ms | ✓ 1624ms | ✓ 959ms | http |
| 128.199.116.219:9090 | ✓ 1406ms | 否 | ✓ 1477ms | ✓ 1144ms | 否 | http |
| 128.199.121.61:9090 | ✓ 1319ms | 否 | ✓ 1460ms | ✓ 1975ms | 否 | http |
| 146.190.80.158:9090 | ✓ 1129ms | 否 | ✓ 1248ms | ✓ 1760ms | ✓ 1557ms | http |
| 103.157.117.116:8080 | 否 | 否 | ✓ 1573ms | ✓ 1709ms | ✓ 1679ms | http |

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
