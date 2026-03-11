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

最后更新：2026-03-11 11:35:25 UTC（2026-03-11 19:35:25 UTC+8）

**代理总数：75**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 75 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 75 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 43.162.113.116:3128 | ✓ 389ms | ✓ 1089ms | ✓ 792ms | ✓ 905ms | ✓ 737ms | http |
| 45.136.131.63:8443 | ✓ 713ms | 否 | ✓ 1775ms | ✓ 686ms | ✓ 534ms | http |
| 1.231.81.166:3128 | ✓ 1250ms | ✓ 941ms | ✓ 909ms | ✓ 938ms | ✓ 701ms | http |
| 160.238.65.7:3128 | ✓ 1408ms | 否 | ✓ 827ms | 否 | ✓ 1737ms | http |
| 160.238.65.3:3128 | ✓ 1408ms | 否 | ✓ 824ms | 否 | ✓ 1746ms | http |
| 160.238.65.5:3128 | ✓ 1405ms | 否 | ✓ 827ms | 否 | ✓ 1751ms | http |
| 160.238.65.9:3128 | ✓ 1409ms | 否 | ✓ 828ms | 否 | ✓ 1749ms | http |
| 160.238.65.8:3128 | ✓ 1405ms | ✓ 1965ms | ✓ 863ms | 否 | ✓ 1762ms | http |
| 160.238.65.4:3128 | ✓ 1409ms | 否 | ✓ 823ms | 否 | ✓ 1762ms | http |
| 160.238.65.6:3128 | ✓ 1405ms | ✓ 1964ms | ✓ 864ms | 否 | ✓ 1746ms | http |
| 211.171.114.154:3128 | ✓ 1471ms | 否 | ✓ 1831ms | ✓ 1474ms | ✓ 1088ms | http |
| 14.143.222.113:10155 | ✓ 1757ms | 否 | ✓ 1117ms | ✓ 1384ms | 否 | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1365ms | ✓ 1227ms | ✓ 975ms | http |
| 45.136.131.47:8443 | ✓ 384ms | ✓ 985ms | ✓ 94ms | ✓ 1936ms | ✓ 507ms | http |
| 103.84.95.54:7890 | ✓ 1083ms | 否 | ✓ 1700ms | ✓ 1231ms | ✓ 957ms | http |
| 95.3.9.78:3128 | ✓ 1158ms | 否 | ✓ 841ms | ✓ 1840ms | ✓ 1403ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1059ms | ✓ 1034ms | 否 | ✓ 978ms | http |
| 175.215.73.250:3180 | 否 | ✓ 1872ms | ✓ 1204ms | ✓ 1155ms | ✓ 980ms | http |
| 35.225.22.61:80 | ✓ 454ms | 否 | ✓ 1441ms | ✓ 1180ms | ✓ 972ms | http |
| 14.225.222.213:7890 | 否 | 否 | ✓ 848ms | ✓ 1695ms | ✓ 1120ms | http |
| 14.225.212.37:7890 | ✓ 1588ms | 否 | ✓ 1213ms | 否 | ✓ 849ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1571ms | 否 | ✓ 1491ms | ✓ 1211ms | http |
| 39.104.201.40:7890 | ✓ 982ms | ✓ 1189ms | ✓ 1493ms | ✓ 1221ms | ✓ 1033ms | http |
| 81.70.169.194:80 | ✓ 1935ms | ✓ 1418ms | ✓ 1221ms | 否 | ✓ 1020ms | http |
| 101.43.255.96:80 | 否 | ✓ 1398ms | ✓ 962ms | ✓ 1307ms | ✓ 1998ms | http |
| 190.9.109.198:999 | ✓ 782ms | 否 | 否 | ✓ 1527ms | ✓ 1201ms | http |
| 46.183.25.8:443 | ✓ 795ms | 否 | ✓ 583ms | ✓ 989ms | ✓ 1296ms | http |
| 45.136.130.175:8443 | ✓ 395ms | ✓ 1651ms | ✓ 241ms | ✓ 1870ms | ✓ 517ms | http |
| 168.235.110.63:3128 | ✓ 497ms | 否 | ✓ 1051ms | ✓ 1425ms | ✓ 1072ms | http |
| 45.136.130.188:8443 | 否 | 否 | ✓ 1351ms | ✓ 839ms | ✓ 702ms | http |
| 194.213.18.200:443 | ✓ 505ms | 否 | 否 | ✓ 1226ms | ✓ 1099ms | http |
| 190.6.54.12:6969 | ✓ 1101ms | ✓ 1742ms | 否 | ✓ 1775ms | 否 | http |
| 202.155.12.161:443 | ✓ 1848ms | 否 | 否 | ✓ 1976ms | ✓ 1573ms | http |
| 45.136.130.223:8443 | ✓ 381ms | ✓ 1175ms | ✓ 91ms | ✓ 688ms | ✓ 525ms | http |
| 45.136.130.191:8443 | ✓ 363ms | 否 | ✓ 82ms | ✓ 661ms | ✓ 519ms | http |
| 42.96.16.158:1311 | ✓ 1361ms | 否 | ✓ 1398ms | ✓ 1138ms | ✓ 920ms | http |
| 45.186.6.104:3128 | ✓ 1381ms | ✓ 1984ms | ✓ 1688ms | 否 | 否 | http |
| 152.70.98.46:8888 | ✓ 1085ms | ✓ 1147ms | ✓ 1780ms | ✓ 1040ms | ✓ 832ms | http |
| 160.238.65.2:3128 | ✓ 1868ms | 否 | ✓ 1552ms | 否 | ✓ 1387ms | http |
| 106.14.203.63:3333 | ✓ 1858ms | 否 | ✓ 950ms | ✓ 1087ms | 否 | http |
| 152.42.213.210:8080 | ✓ 852ms | 否 | ✓ 1099ms | ✓ 1075ms | ✓ 872ms | http |
| 91.107.141.42:8081 | ✓ 1741ms | 否 | ✓ 1385ms | 否 | ✓ 1936ms | http |
| 80.242.56.115:3128 | 否 | ✓ 1911ms | ✓ 1411ms | 否 | ✓ 1619ms | http |
| 113.177.131.2:3128 | ✓ 1965ms | 否 | ✓ 1090ms | ✓ 1339ms | 否 | http |
| 164.90.151.28:3128 | 否 | 否 | ✓ 1033ms | ✓ 994ms | ✓ 1195ms | http |
| 210.223.44.230:3128 | ✓ 1100ms | ✓ 1508ms | ✓ 669ms | ✓ 1246ms | ✓ 1205ms | http |
| 158.69.185.37:3129 | ✓ 1290ms | 否 | ✓ 1949ms | 否 | ✓ 1515ms | http |
| 45.140.147.155:1081 | ✓ 622ms | 否 | ✓ 1396ms | ✓ 1737ms | ✓ 1127ms | http |
| 59.46.216.131:30001 | ✓ 1022ms | 否 | ✓ 1144ms | ✓ 1383ms | 否 | http |
| 205.209.118.30:3138 | ✓ 369ms | 否 | ✓ 1190ms | ✓ 1313ms | ✓ 1007ms | http |
| 45.136.198.40:3128 | ✓ 805ms | 否 | ✓ 775ms | ✓ 1695ms | ✓ 1277ms | http |
| 178.236.245.17:3128 | ✓ 805ms | 否 | ✓ 1557ms | 否 | ✓ 1770ms | http |
| 178.236.245.59:3128 | ✓ 792ms | 否 | ✓ 1565ms | 否 | ✓ 1792ms | http |
| 202.129.206.239:3128 | 否 | 否 | ✓ 1666ms | ✓ 1794ms | ✓ 1774ms | http |
| 103.52.114.95:3128 | ✓ 1821ms | 否 | ✓ 1067ms | ✓ 1249ms | ✓ 982ms | http |
| 8.222.175.80:6128 | ✓ 1499ms | 否 | ✓ 1099ms | ✓ 1060ms | ✓ 935ms | http |
| 190.212.131.238:3128 | ✓ 1148ms | 否 | ✓ 1801ms | 否 | ✓ 1733ms | http |
| 165.227.5.10:8888 | ✓ 563ms | ✓ 1749ms | ✓ 613ms | 否 | ✓ 564ms | http |
| 121.126.185.63:25152 | ✓ 1322ms | 否 | ✓ 1898ms | 否 | ✓ 1655ms | http |
| 34.101.184.164:3128 | ✓ 1638ms | 否 | ✓ 1504ms | ✓ 1305ms | ✓ 1529ms | http |
| 103.183.10.172:3125 | 否 | 否 | ✓ 1792ms | ✓ 1326ms | ✓ 1405ms | http |
| 185.191.236.162:3128 | ✓ 730ms | 否 | ✓ 712ms | ✓ 1934ms | 否 | http |
| 103.183.10.203:3125 | ✓ 1326ms | 否 | ✓ 1310ms | ✓ 1450ms | ✓ 1358ms | http |
| 61.155.242.150:5566 | ✓ 1834ms | 否 | ✓ 1231ms | ✓ 1161ms | 否 | http |
| 47.101.149.27:9010 | 否 | ✓ 1281ms | 否 | ✓ 1441ms | ✓ 1298ms | http |
| 103.39.51.190:8080 | ✓ 1842ms | 否 | 否 | ✓ 1512ms | ✓ 1417ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1322ms | ✓ 1997ms | ✓ 1942ms | http |
| 165.225.72.38:10801 | 否 | 否 | ✓ 1666ms | ✓ 1416ms | ✓ 1247ms | http |
| 165.225.72.38:11526 | 否 | 否 | ✓ 559ms | ✓ 1575ms | ✓ 1237ms | http |
| 165.225.72.38:11670 | 否 | 否 | ✓ 560ms | ✓ 1573ms | ✓ 1247ms | http |
| 103.67.46.225:3125 | 否 | 否 | ✓ 1879ms | ✓ 1666ms | ✓ 1705ms | http |
| 61.52.131.172:8443 | 否 | ✓ 1094ms | ✓ 933ms | 否 | ✓ 941ms | http |
| 45.140.147.155:1082 | ✓ 659ms | ✓ 1590ms | ✓ 1579ms | ✓ 1380ms | ✓ 1192ms | http |
| 45.136.130.239:8443 | 否 | ✓ 1864ms | 否 | ✓ 1468ms | ✓ 1350ms | http |
| 47.74.226.8:5001 | 否 | 否 | ✓ 957ms | ✓ 1462ms | ✓ 1441ms | http |

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
