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

最后更新：2026-03-11 17:49:02 UTC（2026-03-12 01:49:02 UTC+8）

**代理总数：77**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1785ms | ✓ 1324ms | ✓ 1181ms | ✓ 1280ms | ✓ 985ms | http |
| 171.251.172.78:5109 | ✓ 1989ms | 否 | ✓ 1833ms | ✓ 1822ms | ✓ 1609ms | http |
| 35.225.22.61:80 | ✓ 518ms | ✓ 1221ms | ✓ 1047ms | 否 | 否 | http |
| 45.136.130.175:8443 | ✓ 962ms | ✓ 874ms | ✓ 542ms | ✓ 902ms | ✓ 679ms | http |
| 190.9.109.198:999 | ✓ 1105ms | ✓ 1141ms | ✓ 1375ms | ✓ 1463ms | ✓ 1457ms | http |
| 103.84.95.54:7890 | ✓ 1704ms | 否 | 否 | ✓ 1060ms | ✓ 1694ms | http |
| 205.209.118.30:3138 | ✓ 539ms | ✓ 1652ms | ✓ 315ms | ✓ 1068ms | ✓ 963ms | http |
| 45.136.131.63:8443 | ✓ 1072ms | ✓ 936ms | ✓ 1432ms | ✓ 972ms | ✓ 695ms | http |
| 39.104.201.40:7890 | 否 | 否 | ✓ 1070ms | ✓ 1389ms | ✓ 1976ms | http |
| 194.213.18.200:443 | ✓ 522ms | 否 | ✓ 335ms | ✓ 1892ms | 否 | http |
| 152.42.213.210:8080 | ✓ 1263ms | 否 | ✓ 1389ms | ✓ 1584ms | ✓ 1058ms | http |
| 190.212.131.238:3128 | ✓ 797ms | 否 | ✓ 739ms | ✓ 1845ms | 否 | http |
| 152.42.213.210:443 | ✓ 928ms | 否 | 否 | ✓ 1781ms | ✓ 1362ms | http |
| 45.136.131.47:8443 | 否 | ✓ 819ms | ✓ 256ms | ✓ 1421ms | ✓ 666ms | http |
| 101.43.255.96:80 | ✓ 1163ms | 否 | ✓ 1343ms | 否 | ✓ 1183ms | http |
| 45.136.130.239:8443 | 否 | ✓ 893ms | 否 | ✓ 905ms | ✓ 1355ms | http |
| 120.198.141.75:22222 | ✓ 1097ms | 否 | 否 | ✓ 1389ms | ✓ 1180ms | http |
| 202.155.12.161:443 | ✓ 1927ms | 否 | ✓ 1299ms | ✓ 1258ms | ✓ 1957ms | http |
| 81.70.169.194:80 | ✓ 1199ms | 否 | ✓ 1214ms | 否 | ✓ 1196ms | http |
| 121.230.8.111:1080 | 否 | ✓ 1621ms | ✓ 1408ms | ✓ 1804ms | ✓ 1402ms | http |
| 107.173.0.178:1080 | ✓ 993ms | ✓ 1964ms | ✓ 1409ms | 否 | 否 | http |
| 165.227.5.10:8888 | ✓ 331ms | ✓ 1068ms | ✓ 1832ms | ✓ 1234ms | 否 | http |
| 45.136.130.223:8443 | ✓ 997ms | ✓ 1610ms | ✓ 865ms | 否 | 否 | http |
| 45.136.198.40:3128 | ✓ 1178ms | ✓ 1696ms | ✓ 1445ms | ✓ 1866ms | ✓ 1553ms | http |
| 115.231.181.40:8128 | ✓ 1604ms | ✓ 1367ms | ✓ 1104ms | ✓ 1507ms | 否 | http |
| 103.154.214.50:3128 | ✓ 1335ms | 否 | ✓ 1823ms | ✓ 1462ms | ✓ 1173ms | http |
| 91.107.141.42:8081 | ✓ 1060ms | 否 | 否 | ✓ 1293ms | ✓ 1380ms | http |
| 111.48.191.1:7890 | ✓ 1125ms | ✓ 1126ms | ✓ 927ms | ✓ 1994ms | ✓ 973ms | http |
| 152.70.98.46:8888 | ✓ 850ms | 否 | ✓ 1657ms | ✓ 1203ms | ✓ 1352ms | http |
| 162.248.165.72:1080 | ✓ 1112ms | ✓ 1616ms | ✓ 1506ms | ✓ 1272ms | 否 | http |
| 117.159.239.54:22222 | 否 | 否 | ✓ 1033ms | ✓ 1268ms | ✓ 1001ms | http |
| 117.159.239.42:22222 | ✓ 1021ms | ✓ 1370ms | ✓ 949ms | ✓ 1312ms | ✓ 1022ms | http |
| 177.247.249.5:3128 | ✓ 1320ms | ✓ 1477ms | ✓ 1414ms | ✓ 1534ms | 否 | http |
| 120.240.35.176:22222 | ✓ 1088ms | ✓ 1406ms | ✓ 1168ms | ✓ 1357ms | ✓ 1110ms | http |
| 192.71.213.85:9812 | ✓ 1200ms | 否 | ✓ 1567ms | ✓ 1915ms | 否 | http |
| 45.136.130.191:8443 | ✓ 559ms | ✓ 854ms | ✓ 423ms | ✓ 909ms | ✓ 742ms | http |
| 95.3.9.78:8080 | ✓ 1045ms | ✓ 1790ms | 否 | ✓ 1897ms | 否 | http |
| 95.3.9.78:3128 | ✓ 1046ms | 否 | 否 | ✓ 1683ms | ✓ 1968ms | http |
| 210.223.44.230:3128 | ✓ 759ms | ✓ 1549ms | ✓ 967ms | ✓ 1121ms | ✓ 961ms | http |
| 168.235.110.63:3128 | ✓ 342ms | ✓ 1187ms | ✓ 1175ms | ✓ 1607ms | ✓ 902ms | http |
| 47.101.149.27:9010 | 否 | 否 | ✓ 1563ms | ✓ 1810ms | ✓ 1638ms | http |
| 45.136.130.188:8443 | ✓ 911ms | ✓ 1693ms | ✓ 770ms | ✓ 1512ms | ✓ 1337ms | http |
| 150.249.255.91:3128 | ✓ 1790ms | ✓ 1023ms | ✓ 783ms | ✓ 1180ms | ✓ 920ms | http |
| 138.124.53.25:7443 | 否 | ✓ 1638ms | ✓ 1273ms | ✓ 1529ms | ✓ 1585ms | http |
| 106.117.208.101:7890 | ✓ 1227ms | 否 | ✓ 1331ms | 否 | ✓ 1205ms | http |
| 45.186.6.104:3128 | ✓ 1232ms | ✓ 1888ms | ✓ 1674ms | 否 | 否 | http |
| 107.173.52.58:7890 | ✓ 377ms | ✓ 1986ms | ✓ 776ms | ✓ 1192ms | ✓ 1146ms | http |
| 120.198.141.80:22222 | ✓ 1247ms | ✓ 1724ms | ✓ 1257ms | ✓ 1583ms | ✓ 1211ms | http |
| 113.160.132.26:8080 | ✓ 1886ms | ✓ 1665ms | ✓ 1387ms | ✓ 1436ms | ✓ 1219ms | http |
| 180.127.149.252:1080 | ✓ 1182ms | ✓ 1429ms | ✓ 1203ms | 否 | ✓ 1092ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1315ms | ✓ 1711ms | ✓ 1583ms | http |
| 86.53.183.16:1080 | ✓ 1373ms | ✓ 1908ms | ✓ 1382ms | 否 | 否 | http |
| 42.96.16.158:1311 | ✓ 997ms | 否 | ✓ 1147ms | ✓ 1438ms | ✓ 1371ms | http |
| 88.80.150.82:8080 | ✓ 1724ms | ✓ 1780ms | 否 | 否 | ✓ 1813ms | https |
| 172.212.68.37:3128 | ✓ 557ms | ✓ 1437ms | ✓ 750ms | ✓ 1504ms | ✓ 1271ms | http |
| 43.167.227.161:1080 | ✓ 1352ms | 否 | ✓ 1035ms | ✓ 1117ms | 否 | http |
| 222.184.48.235:22222 | ✓ 1173ms | ✓ 1596ms | ✓ 1065ms | 否 | ✓ 1087ms | http |
| 183.249.5.214:22222 | ✓ 1008ms | ✓ 1343ms | ✓ 869ms | ✓ 1123ms | ✓ 849ms | http |
| 120.240.35.177:22222 | 否 | ✓ 1475ms | ✓ 1226ms | ✓ 1473ms | ✓ 1044ms | http |
| 222.184.48.252:22222 | ✓ 989ms | ✓ 1335ms | ✓ 1067ms | ✓ 1516ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1139ms | ✓ 1692ms | 否 | 否 | ✓ 1303ms | http |
| 91.107.148.58:53967 | ✓ 825ms | 否 | 否 | ✓ 1521ms | ✓ 1536ms | http |
| 222.184.48.242:22222 | ✓ 1139ms | ✓ 1349ms | ✓ 1499ms | ✓ 1508ms | ✓ 1071ms | http |
| 222.184.48.241:22222 | ✓ 1132ms | 否 | ✓ 1228ms | ✓ 1530ms | ✓ 1059ms | http |
| 192.71.213.85:9091 | ✓ 620ms | 否 | ✓ 520ms | ✓ 1487ms | 否 | http |
| 113.59.32.163:22222 | ✓ 1222ms | ✓ 1474ms | ✓ 1224ms | ✓ 1477ms | ✓ 1135ms | http |
| 157.254.37.238:999 | ✓ 823ms | ✓ 1813ms | ✓ 1338ms | ✓ 1594ms | ✓ 1606ms | http |
| 61.52.131.172:8443 | ✓ 1084ms | ✓ 1333ms | ✓ 1090ms | ✓ 1327ms | ✓ 1084ms | http |
| 168.138.202.218:3128 | ✓ 1633ms | 否 | 否 | ✓ 1527ms | ✓ 878ms | http |
| 183.249.5.117:22222 | ✓ 1152ms | ✓ 1162ms | ✓ 884ms | ✓ 1123ms | ✓ 960ms | http |
| 178.236.245.17:3128 | ✓ 1483ms | ✓ 1511ms | ✓ 1831ms | ✓ 1905ms | 否 | http |
| 211.171.114.154:3128 | ✓ 1132ms | 否 | ✓ 1992ms | ✓ 1485ms | 否 | http |
| 178.236.245.59:3128 | ✓ 1508ms | 否 | ✓ 1338ms | ✓ 1890ms | 否 | http |
| 111.79.111.126:3128 | ✓ 1790ms | ✓ 1811ms | ✓ 1900ms | 否 | ✓ 1804ms | http |
| 103.39.51.190:8080 | ✓ 1748ms | 否 | ✓ 1899ms | 否 | ✓ 1592ms | http |
| 14.225.222.164:7890 | ✓ 1525ms | ✓ 1790ms | ✓ 1843ms | 否 | 否 | http |
| 222.184.48.251:22222 | ✓ 1857ms | 否 | 否 | ✓ 1961ms | ✓ 1557ms | http |

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
