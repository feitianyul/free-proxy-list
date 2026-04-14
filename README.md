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

最后更新：2026-04-14 12:01:00 UTC（2026-04-14 20:01:00 UTC+8）

**代理总数：70**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 70 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 70 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1639ms | ✓ 1506ms | ✓ 936ms | ✓ 1198ms | ✓ 872ms | http |
| 147.161.239.240:8800 | ✓ 782ms | ✓ 1883ms | ✓ 1446ms | ✓ 1967ms | ✓ 1615ms | http |
| 113.160.132.26:8080 | ✓ 1818ms | 否 | ✓ 1316ms | ✓ 1186ms | ✓ 981ms | http |
| 101.43.127.100:8877 | ✓ 905ms | 否 | ✓ 1163ms | ✓ 1170ms | ✓ 1623ms | http |
| 212.58.132.5:8888 | ✓ 1208ms | 否 | ✓ 1549ms | ✓ 1624ms | ✓ 1333ms | http |
| 114.118.82.146:80 | ✓ 1058ms | ✓ 1165ms | ✓ 1066ms | ✓ 1171ms | ✓ 899ms | http |
| 5.104.87.17:8051 | ✓ 910ms | 否 | ✓ 712ms | ✓ 893ms | ✓ 1013ms | http |
| 157.230.178.216:8088 | ✓ 627ms | 否 | ✓ 657ms | ✓ 1313ms | ✓ 1081ms | http |
| 8.219.97.248:80 | ✓ 1228ms | 否 | ✓ 843ms | ✓ 1372ms | 否 | http |
| 34.71.229.255:3128 | ✓ 1022ms | 否 | ✓ 1442ms | ✓ 1206ms | ✓ 1248ms | http |
| 167.103.115.102:8800 | ✓ 1182ms | 否 | ✓ 1183ms | 否 | ✓ 959ms | http |
| 167.103.144.127:8800 | ✓ 1481ms | 否 | ✓ 1760ms | ✓ 1693ms | ✓ 1571ms | http |
| 85.239.59.252:7890 | ✓ 851ms | 否 | ✓ 1572ms | 否 | ✓ 1550ms | http |
| 167.103.34.108:8800 | ✓ 1720ms | 否 | ✓ 1358ms | 否 | ✓ 1591ms | http |
| 94.131.118.39:1081 | ✓ 1218ms | 否 | ✓ 1304ms | ✓ 1807ms | 否 | http |
| 78.11.96.22:8888 | ✓ 1514ms | 否 | ✓ 1503ms | ✓ 1737ms | ✓ 1573ms | http |
| 45.167.125.21:999 | ✓ 1319ms | 否 | ✓ 1490ms | ✓ 1844ms | ✓ 1627ms | http |
| 167.103.31.122:8800 | ✓ 1611ms | 否 | ✓ 1320ms | ✓ 1662ms | 否 | http |
| 120.92.108.86:7890 | ✓ 1857ms | 否 | ✓ 1219ms | 否 | ✓ 1541ms | http |
| 43.99.54.80:7890 | 否 | 否 | ✓ 633ms | ✓ 1695ms | ✓ 682ms | http |
| 35.225.22.61:80 | ✓ 753ms | ✓ 1493ms | ✓ 934ms | 否 | 否 | http |
| 52.16.215.4:30568 | ✓ 1662ms | 否 | ✓ 800ms | 否 | ✓ 1723ms | http |
| 77.110.113.24:40000 | ✓ 1110ms | 否 | ✓ 987ms | 否 | ✓ 1778ms | http |
| 72.56.84.21:1080 | 否 | 否 | ✓ 1343ms | ✓ 1577ms | ✓ 1642ms | http |
| 138.124.99.216:8888 | ✓ 815ms | ✓ 1856ms | 否 | 否 | ✓ 1712ms | http |
| 82.114.228.67:1080 | ✓ 1481ms | ✓ 1675ms | ✓ 1515ms | 否 | 否 | http |
| 94.131.118.129:1081 | ✓ 1208ms | ✓ 1892ms | ✓ 883ms | ✓ 1702ms | ✓ 1270ms | http |
| 2.27.18.184:1080 | ✓ 894ms | ✓ 1584ms | ✓ 1681ms | 否 | 否 | http |
| 1.231.81.166:3128 | ✓ 1655ms | ✓ 1129ms | ✓ 840ms | 否 | 否 | http |
| 181.78.44.63:999 | ✓ 1201ms | 否 | ✓ 1352ms | ✓ 1786ms | ✓ 1490ms | http |
| 133.18.110.87:1081 | ✓ 1414ms | ✓ 1153ms | ✓ 791ms | ✓ 816ms | ✓ 808ms | http |
| 185.114.73.2:1080 | ✓ 998ms | ✓ 1537ms | ✓ 1822ms | 否 | 否 | http |
| 5.255.123.43:1080 | ✓ 619ms | 否 | ✓ 1782ms | ✓ 1743ms | ✓ 1561ms | http |
| 115.231.181.40:8128 | ✓ 1971ms | 否 | ✓ 1681ms | ✓ 1118ms | ✓ 931ms | http |
| 116.80.63.178:3172 | ✓ 1551ms | 否 | ✓ 1545ms | ✓ 1774ms | 否 | http |
| 34.101.184.164:3128 | ✓ 858ms | 否 | ✓ 1613ms | 否 | ✓ 1857ms | http |
| 144.31.27.49:1080 | ✓ 1292ms | 否 | ✓ 1967ms | 否 | ✓ 1942ms | http |
| 121.230.9.248:1080 | ✓ 1088ms | ✓ 1489ms | ✓ 1235ms | 否 | ✓ 1239ms | http |
| 121.230.8.201:1080 | 否 | ✓ 1228ms | 否 | ✓ 1519ms | ✓ 1065ms | http |
| 47.84.131.156:8100 | ✓ 712ms | 否 | ✓ 704ms | ✓ 1017ms | ✓ 829ms | http |
| 152.32.132.190:7890 | 否 | 否 | ✓ 1635ms | ✓ 1444ms | ✓ 821ms | http |
| 3.96.208.241:22894 | ✓ 1319ms | 否 | ✓ 1909ms | 否 | ✓ 1889ms | http |
| 12.89.176.82:3128 | ✓ 719ms | ✓ 1339ms | ✓ 1258ms | ✓ 1887ms | ✓ 1940ms | http |
| 101.32.243.189:80 | ✓ 1060ms | ✓ 1494ms | ✓ 1477ms | ✓ 1371ms | ✓ 1813ms | http |
| 171.244.130.36:3128 | ✓ 1561ms | 否 | ✓ 1371ms | ✓ 1817ms | ✓ 1335ms | http |
| 59.46.216.131:30001 | ✓ 1935ms | 否 | ✓ 1617ms | 否 | ✓ 1087ms | http |
| 177.234.217.88:999 | ✓ 1085ms | ✓ 1881ms | ✓ 1034ms | ✓ 1908ms | ✓ 1595ms | http |
| 107.172.102.234:40621 | 否 | ✓ 1809ms | ✓ 1161ms | ✓ 867ms | ✓ 1051ms | http |
| 223.84.151.86:30005 | ✓ 1575ms | ✓ 1819ms | ✓ 1275ms | ✓ 1854ms | ✓ 1988ms | http |
| 210.77.22.250:7890 | ✓ 1089ms | ✓ 1211ms | ✓ 1809ms | 否 | ✓ 889ms | http |
| 36.141.21.200:7890 | 否 | 否 | ✓ 848ms | ✓ 1292ms | ✓ 948ms | http |
| 94.72.109.214:8888 | ✓ 1460ms | 否 | ✓ 1974ms | 否 | ✓ 1752ms | http |
| 5.104.87.17:8050 | ✓ 1704ms | 否 | ✓ 1534ms | 否 | ✓ 1474ms | http |
| 34.246.183.20:8081 | 否 | 否 | ✓ 1922ms | ✓ 1822ms | ✓ 1360ms | http |
| 3.71.26.7:40592 | ✓ 851ms | 否 | ✓ 1690ms | 否 | ✓ 1626ms | http |
| 94.23.12.224:8888 | ✓ 1754ms | 否 | ✓ 1688ms | 否 | ✓ 1717ms | http |
| 217.217.249.160:8080 | ✓ 1524ms | 否 | ✓ 1477ms | 否 | ✓ 1830ms | http |
| 157.0.142.246:10061 | ✓ 961ms | ✓ 1179ms | ✓ 951ms | ✓ 1260ms | ✓ 970ms | http |
| 45.140.147.82:1081 | ✓ 1225ms | ✓ 1485ms | ✓ 1715ms | ✓ 1804ms | ✓ 1154ms | http |
| 210.223.44.230:3128 | ✓ 1101ms | ✓ 1873ms | ✓ 672ms | 否 | ✓ 1147ms | http |
| 171.227.167.109:1004 | ✓ 856ms | 否 | ✓ 983ms | ✓ 1192ms | ✓ 995ms | http |
| 116.80.83.23:3128 | ✓ 1550ms | 否 | 否 | ✓ 1901ms | ✓ 1754ms | http |
| 223.16.170.103:80 | ✓ 1627ms | 否 | ✓ 945ms | ✓ 1026ms | ✓ 1033ms | http |
| 144.31.140.92:1080 | ✓ 734ms | 否 | ✓ 851ms | 否 | ✓ 1902ms | http |
| 61.52.131.172:8443 | ✓ 865ms | ✓ 1168ms | ✓ 919ms | ✓ 1212ms | ✓ 945ms | http |
| 116.80.59.73:3128 | ✓ 1578ms | 否 | 否 | ✓ 1745ms | ✓ 1646ms | http |
| 185.191.236.162:3128 | ✓ 1755ms | ✓ 1812ms | 否 | 否 | ✓ 1523ms | http |
| 104.129.203.245:10733 | ✓ 1069ms | ✓ 666ms | ✓ 594ms | ✓ 671ms | ✓ 549ms | http |
| 104.129.203.245:10139 | ✓ 1214ms | ✓ 585ms | ✓ 540ms | ✓ 875ms | ✓ 529ms | http |
| 104.129.203.245:10026 | ✓ 1105ms | ✓ 1145ms | ✓ 572ms | ✓ 657ms | ✓ 488ms | http |

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
