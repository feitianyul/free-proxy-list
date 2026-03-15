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

最后更新：2026-03-15 09:31:52 UTC（2026-03-15 17:31:52 UTC+8）

**代理总数：55**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 54 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 55 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 257ms | 否 | ✓ 632ms | ✓ 930ms | ✓ 765ms | http |
| 38.145.218.82:8443 | 否 | 否 | ✓ 501ms | ✓ 977ms | ✓ 738ms | http |
| 113.160.132.26:8080 | ✓ 1944ms | 否 | ✓ 1345ms | ✓ 1514ms | ✓ 1189ms | http |
| 45.167.124.52:8080 | ✓ 1048ms | 否 | ✓ 1205ms | ✓ 1598ms | ✓ 1306ms | http |
| 47.101.149.27:9010 | 否 | 否 | ✓ 1508ms | ✓ 1760ms | ✓ 1582ms | http |
| 120.92.212.16:7890 | ✓ 1122ms | ✓ 1403ms | 否 | 否 | ✓ 1129ms | http |
| 217.76.245.80:999 | ✓ 647ms | 否 | ✓ 1542ms | ✓ 1411ms | 否 | http |
| 35.225.22.61:80 | ✓ 1228ms | ✓ 1926ms | ✓ 862ms | 否 | 否 | http |
| 160.250.223.216:10216 | ✓ 1622ms | 否 | ✓ 1850ms | 否 | ✓ 1899ms | http |
| 120.92.212.16:8890 | ✓ 1128ms | 否 | ✓ 1391ms | ✓ 1742ms | ✓ 1167ms | http |
| 101.43.127.100:8877 | ✓ 1120ms | ✓ 1234ms | ✓ 1494ms | ✓ 1418ms | ✓ 1063ms | http |
| 38.55.105.94:6005 | ✓ 1369ms | 否 | ✓ 1504ms | ✓ 1493ms | ✓ 1003ms | http |
| 137.220.150.152:6005 | ✓ 1260ms | 否 | ✓ 1780ms | ✓ 1322ms | ✓ 1056ms | http |
| 185.115.74.185:8080 | ✓ 802ms | ✓ 1848ms | ✓ 1847ms | 否 | 否 | http |
| 165.227.5.10:8888 | ✓ 332ms | ✓ 1162ms | ✓ 1439ms | ✓ 1370ms | ✓ 1591ms | http |
| 101.47.73.135:3128 | ✓ 1689ms | 否 | ✓ 886ms | ✓ 1676ms | ✓ 1324ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1185ms | ✓ 1366ms | ✓ 1043ms | http |
| 81.70.169.194:80 | ✓ 1203ms | ✓ 1501ms | 否 | ✓ 1479ms | ✓ 1884ms | http |
| 95.85.252.153:21064 | ✓ 974ms | ✓ 1718ms | ✓ 1524ms | ✓ 1865ms | ✓ 1363ms | http |
| 137.220.151.110:6005 | ✓ 919ms | 否 | ✓ 1233ms | ✓ 1354ms | ✓ 1067ms | http |
| 137.220.150.104:6005 | ✓ 910ms | 否 | ✓ 1241ms | ✓ 1367ms | ✓ 1111ms | http |
| 85.198.96.242:3128 | ✓ 1616ms | ✓ 1772ms | 否 | 否 | ✓ 1197ms | http |
| 38.55.106.208:6005 | ✓ 1704ms | 否 | ✓ 1140ms | 否 | ✓ 1142ms | http |
| 101.43.255.96:80 | ✓ 1457ms | ✓ 1757ms | ✓ 1087ms | 否 | ✓ 1567ms | http |
| 212.192.13.76:6005 | ✓ 1773ms | 否 | ✓ 1827ms | 否 | ✓ 1848ms | http |
| 140.238.156.12:1080 | ✓ 403ms | ✓ 1858ms | 否 | ✓ 1763ms | 否 | http |
| 62.60.177.204:34094 | ✓ 374ms | 否 | ✓ 599ms | ✓ 1053ms | ✓ 811ms | http |
| 103.84.95.54:7890 | ✓ 835ms | 否 | 否 | ✓ 1028ms | ✓ 819ms | http |
| 172.212.68.37:3128 | ✓ 709ms | ✓ 1595ms | ✓ 1206ms | ✓ 1164ms | ✓ 1011ms | http |
| 38.145.218.101:8447 | ✓ 1131ms | ✓ 1770ms | ✓ 470ms | ✓ 961ms | ✓ 772ms | http |
| 38.145.218.235:8443 | ✓ 1396ms | ✓ 1076ms | ✓ 294ms | ✓ 904ms | ✓ 1411ms | http |
| 38.145.218.161:8443 | ✓ 1398ms | 否 | ✓ 281ms | ✓ 950ms | ✓ 739ms | http |
| 88.80.150.82:8080 | ✓ 1174ms | 否 | ✓ 1046ms | ✓ 1903ms | ✓ 1592ms | https |
| 101.32.244.83:8080 | ✓ 1223ms | ✓ 1844ms | ✓ 1123ms | ✓ 1538ms | ✓ 1520ms | http |
| 121.43.196.213:8222 | ✓ 1106ms | ✓ 1304ms | ✓ 1027ms | ✓ 1371ms | ✓ 1053ms | http |
| 121.43.196.210:8222 | ✓ 1121ms | ✓ 1255ms | ✓ 1081ms | ✓ 1441ms | ✓ 1041ms | http |
| 114.55.226.123:10086 | ✓ 1209ms | ✓ 1916ms | ✓ 1172ms | ✓ 1490ms | ✓ 1194ms | http |
| 92.119.127.213:6005 | 否 | ✓ 1756ms | ✓ 1500ms | ✓ 1731ms | ✓ 1422ms | http |
| 45.136.198.40:3128 | ✓ 1044ms | 否 | ✓ 1763ms | 否 | ✓ 1474ms | http |
| 104.129.202.127:12354 | ✓ 1346ms | ✓ 1915ms | ✓ 1221ms | ✓ 1607ms | ✓ 1281ms | http |
| 104.129.202.127:10810 | ✓ 1345ms | 否 | ✓ 1143ms | ✓ 1623ms | ✓ 1284ms | http |
| 38.180.2.107:3128 | ✓ 1287ms | 否 | ✓ 1988ms | 否 | ✓ 1784ms | http |
| 86.53.183.16:1080 | ✓ 433ms | ✓ 1658ms | ✓ 1741ms | 否 | 否 | http |
| 45.119.85.216:3128 | ✓ 1575ms | 否 | ✓ 1702ms | ✓ 1577ms | ✓ 1231ms | http |
| 207.254.71.62:8088 | ✓ 1897ms | 否 | ✓ 1467ms | ✓ 1545ms | ✓ 1505ms | http |
| 38.145.203.135:8443 | ✓ 734ms | ✓ 1624ms | ✓ 792ms | ✓ 1201ms | ✓ 731ms | http |
| 103.139.138.194:3128 | ✓ 1984ms | 否 | ✓ 1709ms | ✓ 1708ms | ✓ 1419ms | http |
| 121.40.231.103:7890 | 否 | ✓ 1807ms | ✓ 1705ms | 否 | ✓ 1868ms | http |
| 61.52.131.172:8443 | ✓ 1077ms | ✓ 1336ms | ✓ 1068ms | ✓ 1375ms | ✓ 1117ms | http |
| 36.155.100.217:8080 | ✓ 1514ms | ✓ 1411ms | ✓ 1293ms | ✓ 1686ms | ✓ 1415ms | http |
| 45.136.131.42:8447 | ✓ 1936ms | ✓ 1357ms | ✓ 1010ms | ✓ 1353ms | ✓ 1425ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1335ms | ✓ 1648ms | ✓ 1293ms | http |
| 45.149.92.147:5001 | ✓ 796ms | 否 | ✓ 808ms | ✓ 1004ms | ✓ 813ms | http |
| 124.16.111.161:7890 | ✓ 1018ms | ✓ 1319ms | ✓ 1224ms | ✓ 1301ms | ✓ 1033ms | http |
| 47.77.193.180:1080 | ✓ 903ms | 否 | ✓ 513ms | ✓ 989ms | ✓ 723ms | http |

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
