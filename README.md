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

最后更新：2026-02-28 11:24:10 UTC（2026-02-28 19:24:10 UTC+8）

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
| 205.209.118.30:3138 | ✓ 531ms | ✓ 1288ms | ✓ 1134ms | ✓ 1336ms | ✓ 1066ms | http |
| 3.213.157.4:3128 | ✓ 328ms | 否 | ✓ 283ms | ✓ 1200ms | ✓ 925ms | http |
| 195.123.209.48:3128 | ✓ 1394ms | 否 | ✓ 1798ms | 否 | ✓ 1527ms | http |
| 35.234.17.221:8080 | 否 | ✓ 1291ms | ✓ 1037ms | ✓ 960ms | ✓ 1134ms | http |
| 138.124.53.25:7443 | ✓ 826ms | 否 | ✓ 1559ms | 否 | ✓ 1768ms | http |
| 120.92.212.16:8890 | ✓ 921ms | ✓ 1162ms | ✓ 1021ms | ✓ 1184ms | ✓ 942ms | http |
| 120.92.212.16:7890 | ✓ 913ms | ✓ 1187ms | ✓ 1028ms | ✓ 1411ms | ✓ 960ms | http |
| 81.70.169.194:80 | ✓ 973ms | 否 | ✓ 1128ms | ✓ 1130ms | ✓ 1023ms | http |
| 103.84.95.54:7890 | ✓ 666ms | 否 | ✓ 917ms | ✓ 881ms | 否 | http |
| 132.145.93.138:1080 | ✓ 1540ms | 否 | 否 | ✓ 1666ms | ✓ 1136ms | http |
| 36.147.78.166:80 | 否 | ✓ 1533ms | 否 | ✓ 1720ms | ✓ 1609ms | http |
| 101.43.255.96:80 | ✓ 970ms | ✓ 1172ms | ✓ 955ms | ✓ 1251ms | 否 | http |
| 62.113.119.14:8080 | ✓ 992ms | 否 | ✓ 1560ms | 否 | ✓ 1266ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1347ms | ✓ 1420ms | ✓ 1051ms | http |
| 91.233.223.147:3128 | ✓ 1062ms | 否 | ✓ 1893ms | 否 | ✓ 1727ms | http |
| 168.235.110.63:3128 | ✓ 445ms | 否 | ✓ 1384ms | 否 | ✓ 1880ms | http |
| 121.237.181.137:8888 | 否 | ✓ 1159ms | ✓ 899ms | 否 | ✓ 784ms | http |
| 139.159.97.82:10900 | ✓ 1056ms | 否 | ✓ 1672ms | ✓ 1703ms | ✓ 1191ms | http |
| 115.231.181.40:8128 | ✓ 823ms | ✓ 1832ms | 否 | ✓ 1116ms | 否 | http |
| 111.79.111.126:3128 | 否 | 否 | ✓ 1868ms | ✓ 1870ms | ✓ 1431ms | http |
| 190.9.109.202:999 | ✓ 1159ms | ✓ 1395ms | ✓ 1345ms | 否 | 否 | http |
| 167.103.6.46:11197 | 否 | 否 | ✓ 1270ms | ✓ 1789ms | ✓ 1236ms | http |
| 35.225.22.61:80 | ✓ 437ms | ✓ 1357ms | 否 | ✓ 1001ms | 否 | http |
| 142.171.85.32:1080 | ✓ 325ms | 否 | ✓ 1737ms | ✓ 688ms | ✓ 1033ms | http |
| 90.84.188.97:8000 | ✓ 1704ms | 否 | ✓ 1757ms | ✓ 1721ms | ✓ 1864ms | http |
| 45.125.67.37:8443 | ✓ 960ms | 否 | ✓ 1813ms | ✓ 1011ms | ✓ 971ms | http |
| 165.227.5.10:8888 | ✓ 565ms | 否 | ✓ 431ms | ✓ 1735ms | 否 | http |
| 81.177.48.54:2080 | ✓ 1005ms | 否 | ✓ 1972ms | 否 | ✓ 1507ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1090ms | 否 | ✓ 1462ms | ✓ 971ms | http |
| 83.219.250.8:62920 | ✓ 976ms | 否 | ✓ 1417ms | 否 | ✓ 1857ms | http |
| 120.132.97.88:7897 | ✓ 865ms | 否 | ✓ 912ms | ✓ 1114ms | ✓ 873ms | http |
| 165.225.122.12:11685 | ✓ 1905ms | 否 | ✓ 1344ms | 否 | ✓ 1537ms | http |
| 136.226.254.24:12360 | 否 | ✓ 1976ms | ✓ 1517ms | ✓ 1854ms | ✓ 1561ms | http |
| 136.226.254.24:11909 | 否 | 否 | ✓ 1550ms | ✓ 1834ms | ✓ 1546ms | http |
| 136.226.254.24:11819 | ✓ 1893ms | 否 | 否 | ✓ 1560ms | ✓ 1539ms | http |
| 167.103.6.24:11836 | ✓ 1906ms | 否 | ✓ 1727ms | ✓ 1534ms | 否 | http |
| 136.226.254.24:12265 | ✓ 1900ms | ✓ 1996ms | ✓ 1603ms | ✓ 1978ms | ✓ 1540ms | http |
| 136.226.254.24:11933 | ✓ 1899ms | 否 | 否 | ✓ 1607ms | ✓ 1555ms | http |
| 136.226.254.24:11342 | ✓ 1900ms | 否 | 否 | ✓ 1613ms | ✓ 1560ms | http |
| 147.45.216.148:1080 | ✓ 1276ms | 否 | ✓ 1122ms | ✓ 1468ms | ✓ 1018ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1087ms | ✓ 1100ms | ✓ 995ms | http |
| 121.230.8.235:1080 | ✓ 1036ms | ✓ 1654ms | ✓ 1340ms | ✓ 1235ms | ✓ 1046ms | http |
| 121.230.9.37:1080 | ✓ 1084ms | 否 | ✓ 1243ms | ✓ 1743ms | ✓ 1410ms | http |
| 121.230.9.148:1080 | 否 | 否 | ✓ 1375ms | ✓ 1459ms | ✓ 999ms | http |
| 101.32.244.83:8080 | ✓ 1479ms | ✓ 1304ms | ✓ 880ms | ✓ 1362ms | ✓ 1156ms | http |
| 121.43.196.210:8222 | ✓ 878ms | ✓ 1041ms | ✓ 861ms | ✓ 1080ms | ✓ 808ms | http |
| 121.43.196.213:8222 | ✓ 932ms | ✓ 988ms | ✓ 869ms | ✓ 1072ms | ✓ 817ms | http |
| 114.55.226.123:10086 | ✓ 1023ms | ✓ 1855ms | ✓ 937ms | ✓ 1239ms | ✓ 1000ms | http |
| 43.161.214.161:1081 | ✓ 1741ms | 否 | 否 | ✓ 1902ms | ✓ 1361ms | http |
| 45.140.147.82:1081 | ✓ 1145ms | 否 | ✓ 1543ms | ✓ 1697ms | ✓ 1260ms | http |
| 165.225.122.12:12125 | ✓ 1240ms | 否 | ✓ 1237ms | ✓ 1568ms | 否 | http |
| 165.225.122.12:11691 | ✓ 1224ms | 否 | ✓ 1243ms | ✓ 1572ms | ✓ 1426ms | http |
| 165.225.122.12:10919 | ✓ 1229ms | 否 | ✓ 1291ms | ✓ 1685ms | ✓ 1298ms | http |
| 165.225.122.12:12302 | ✓ 1234ms | 否 | ✓ 1238ms | 否 | ✓ 1131ms | http |
| 167.103.6.24:12174 | ✓ 1237ms | 否 | ✓ 1291ms | ✓ 1733ms | ✓ 1398ms | http |
| 167.103.6.24:12363 | ✓ 1230ms | 否 | ✓ 1774ms | ✓ 1538ms | ✓ 1180ms | http |
| 167.103.6.24:11016 | ✓ 1226ms | 否 | ✓ 1289ms | 否 | ✓ 1204ms | http |
| 165.225.122.17:11146 | ✓ 1221ms | 否 | ✓ 1322ms | 否 | ✓ 1193ms | http |
| 167.103.6.24:11263 | ✓ 1234ms | 否 | ✓ 1340ms | ✓ 1516ms | ✓ 1707ms | http |
| 167.103.6.24:11974 | ✓ 1234ms | 否 | ✓ 1242ms | 否 | ✓ 1399ms | http |
| 165.225.122.17:10122 | ✓ 1222ms | 否 | ✓ 1241ms | ✓ 1544ms | 否 | http |
| 165.225.122.12:11962 | ✓ 1225ms | 否 | ✓ 1241ms | ✓ 1635ms | 否 | http |
| 45.140.147.155:1081 | ✓ 1768ms | 否 | ✓ 1447ms | ✓ 1750ms | ✓ 1095ms | http |
| 44.205.216.127:80 | ✓ 491ms | ✓ 1342ms | ✓ 1087ms | ✓ 1177ms | ✓ 903ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 991ms | ✓ 881ms | ✓ 762ms | http |
| 8.219.97.248:80 | ✓ 1447ms | 否 | ✓ 1051ms | 否 | ✓ 1156ms | http |
| 36.147.78.166:443 | 否 | 否 | ✓ 1562ms | ✓ 1787ms | ✓ 1523ms | http |
| 180.130.80.196:9003 | ✓ 1761ms | ✓ 1356ms | ✓ 1310ms | ✓ 1318ms | ✓ 1352ms | http |
| 212.175.29.184:8080 | ✓ 1422ms | 否 | ✓ 1846ms | 否 | ✓ 1709ms | http |
| 84.247.149.172:3128 | ✓ 1534ms | 否 | ✓ 1658ms | ✓ 1025ms | ✓ 801ms | http |
| 45.136.198.40:3128 | ✓ 853ms | 否 | ✓ 831ms | ✓ 1698ms | ✓ 1343ms | http |
| 45.140.147.82:1082 | ✓ 619ms | 否 | ✓ 1444ms | 否 | ✓ 1330ms | http |
| 121.230.8.211:1080 | ✓ 1171ms | ✓ 1311ms | 否 | ✓ 1631ms | ✓ 1114ms | http |
| 103.39.51.190:8080 | ✓ 1991ms | 否 | 否 | ✓ 1292ms | ✓ 1848ms | http |
| 85.208.108.43:2094 | ✓ 826ms | 否 | ✓ 1112ms | ✓ 1193ms | ✓ 887ms | http |

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
