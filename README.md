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

最后更新：2026-03-16 20:02:05 UTC（2026-03-17 04:02:05 UTC+8）

**代理总数：77**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1534ms | ✓ 928ms | ✓ 1379ms | ✓ 984ms | ✓ 858ms | http |
| 137.220.151.110:6005 | 否 | 否 | ✓ 1014ms | ✓ 1332ms | ✓ 862ms | http |
| 202.155.12.161:443 | ✓ 1504ms | ✓ 1977ms | ✓ 775ms | ✓ 1147ms | ✓ 932ms | http |
| 147.161.210.140:8800 | ✓ 1505ms | 否 | ✓ 844ms | ✓ 998ms | ✓ 1066ms | http |
| 113.160.132.26:8080 | ✓ 1409ms | ✓ 1348ms | ✓ 1280ms | ✓ 1193ms | ✓ 979ms | http |
| 115.231.181.40:8128 | ✓ 913ms | ✓ 1038ms | ✓ 930ms | 否 | 否 | http |
| 149.50.116.240:1080 | ✓ 1008ms | 否 | ✓ 1965ms | ✓ 1949ms | ✓ 1663ms | http |
| 137.220.150.152:6005 | ✓ 1022ms | ✓ 1950ms | ✓ 866ms | 否 | 否 | http |
| 45.167.124.52:8080 | 否 | 否 | ✓ 1439ms | ✓ 1923ms | ✓ 1686ms | http |
| 45.136.198.40:3128 | ✓ 1152ms | ✓ 1948ms | ✓ 1688ms | ✓ 1953ms | ✓ 1944ms | http |
| 8.209.239.31:30000 | ✓ 1023ms | ✓ 1555ms | ✓ 610ms | ✓ 816ms | ✓ 600ms | http |
| 8.219.97.248:80 | ✓ 1783ms | 否 | 否 | ✓ 1893ms | ✓ 1531ms | http |
| 103.84.95.54:7890 | ✓ 1012ms | 否 | ✓ 683ms | ✓ 1190ms | ✓ 784ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1343ms | ✓ 1119ms | ✓ 1376ms | ✓ 1104ms | http |
| 121.230.8.97:1080 | ✓ 1972ms | ✓ 1587ms | ✓ 993ms | ✓ 1483ms | ✓ 1306ms | http |
| 101.43.127.100:8877 | ✓ 953ms | ✓ 1953ms | 否 | ✓ 1152ms | ✓ 1376ms | http |
| 212.192.13.76:6005 | 否 | 否 | ✓ 1289ms | ✓ 1319ms | ✓ 1668ms | http |
| 190.242.157.215:8080 | ✓ 1411ms | ✓ 1937ms | ✓ 1459ms | 否 | 否 | http |
| 62.60.177.204:34094 | ✓ 372ms | 否 | ✓ 842ms | 否 | ✓ 841ms | http |
| 45.186.6.104:3128 | ✓ 1246ms | ✓ 1949ms | ✓ 1941ms | 否 | 否 | http |
| 72.11.151.159:6005 | ✓ 860ms | ✓ 1262ms | ✓ 833ms | ✓ 1479ms | ✓ 1031ms | http |
| 35.225.22.61:80 | ✓ 997ms | 否 | ✓ 1027ms | ✓ 1288ms | ✓ 965ms | http |
| 137.220.150.170:6005 | ✓ 747ms | 否 | ✓ 744ms | ✓ 1141ms | ✓ 896ms | http |
| 38.55.107.137:6005 | ✓ 1205ms | 否 | ✓ 1002ms | ✓ 1142ms | ✓ 680ms | http |
| 212.192.12.90:6005 | ✓ 1205ms | 否 | ✓ 1273ms | ✓ 1303ms | ✓ 729ms | http |
| 178.236.245.59:3128 | ✓ 871ms | 否 | ✓ 934ms | ✓ 1738ms | ✓ 1402ms | http |
| 159.223.42.219:3128 | 否 | 否 | ✓ 916ms | ✓ 1240ms | ✓ 1174ms | http |
| 178.236.245.17:3128 | ✓ 743ms | 否 | ✓ 1070ms | ✓ 1834ms | 否 | http |
| 165.225.72.38:10057 | ✓ 905ms | 否 | ✓ 1195ms | ✓ 1750ms | ✓ 1494ms | http |
| 165.225.72.38:10056 | ✓ 892ms | 否 | ✓ 937ms | ✓ 1932ms | ✓ 1584ms | http |
| 165.225.72.38:10847 | ✓ 599ms | ✓ 1816ms | ✓ 536ms | ✓ 1524ms | ✓ 1442ms | http |
| 165.225.72.38:10960 | ✓ 616ms | 否 | ✓ 527ms | ✓ 1504ms | ✓ 1285ms | http |
| 165.225.72.38:11211 | ✓ 544ms | ✓ 1785ms | ✓ 525ms | ✓ 1517ms | 否 | http |
| 168.235.110.63:3128 | ✓ 834ms | 否 | ✓ 960ms | ✓ 1129ms | ✓ 860ms | http |
| 137.220.150.104:6005 | ✓ 1165ms | 否 | ✓ 1223ms | ✓ 1250ms | ✓ 1133ms | http |
| 120.92.212.16:7890 | ✓ 1025ms | ✓ 1868ms | 否 | ✓ 1509ms | ✓ 971ms | http |
| 114.231.72.27:1080 | 否 | 否 | ✓ 1930ms | ✓ 1183ms | ✓ 976ms | http |
| 165.225.72.38:11515 | ✓ 871ms | ✓ 1808ms | ✓ 1341ms | 否 | 否 | http |
| 120.92.212.16:8890 | ✓ 960ms | ✓ 1234ms | ✓ 952ms | 否 | 否 | http |
| 194.5.212.40:8080 | ✓ 1054ms | 否 | ✓ 1144ms | ✓ 1735ms | ✓ 1816ms | http |
| 165.225.72.38:10801 | ✓ 1041ms | ✓ 1824ms | ✓ 1292ms | ✓ 1763ms | ✓ 1620ms | http |
| 121.40.231.103:7890 | 否 | ✓ 1256ms | ✓ 1886ms | 否 | ✓ 1402ms | http |
| 133.242.138.34:8100 | ✓ 1934ms | 否 | ✓ 1083ms | ✓ 1988ms | ✓ 1695ms | http |
| 116.80.49.167:3172 | ✓ 1522ms | 否 | ✓ 1515ms | ✓ 1875ms | 否 | http |
| 113.255.59.226:8080 | 否 | 否 | ✓ 1097ms | ✓ 1113ms | ✓ 1158ms | http |
| 106.117.208.101:7890 | ✓ 1031ms | ✓ 1586ms | ✓ 1921ms | ✓ 1752ms | ✓ 1025ms | http |
| 165.225.72.38:10561 | ✓ 1041ms | ✓ 1815ms | ✓ 1417ms | ✓ 1965ms | ✓ 1452ms | http |
| 165.225.72.38:33333 | ✓ 1041ms | ✓ 1806ms | ✓ 1431ms | ✓ 1804ms | ✓ 1568ms | http |
| 2.56.122.146:10808 | ✓ 1183ms | 否 | ✓ 1626ms | 否 | ✓ 1280ms | http |
| 219.117.204.211:7799 | ✓ 643ms | ✓ 841ms | ✓ 936ms | ✓ 843ms | ✓ 1171ms | http |
| 172.105.118.164:3128 | 否 | ✓ 1946ms | ✓ 1596ms | 否 | ✓ 1748ms | http |
| 45.136.131.55:8452 | ✓ 731ms | ✓ 1212ms | ✓ 687ms | ✓ 753ms | ✓ 868ms | http |
| 103.39.51.190:8080 | ✓ 1849ms | 否 | 否 | ✓ 1477ms | ✓ 1387ms | http |
| 62.113.119.14:8080 | ✓ 1196ms | ✓ 1888ms | ✓ 1036ms | 否 | ✓ 1393ms | http |
| 165.225.72.38:10639 | ✓ 754ms | 否 | ✓ 1932ms | 否 | ✓ 1608ms | http |
| 165.225.72.38:10011 | ✓ 1059ms | 否 | ✓ 1643ms | 否 | ✓ 1670ms | http |
| 165.225.72.38:11814 | ✓ 821ms | 否 | ✓ 897ms | ✓ 1712ms | ✓ 1260ms | http |
| 165.225.72.38:11112 | ✓ 825ms | 否 | ✓ 873ms | ✓ 1678ms | ✓ 1121ms | http |
| 165.225.72.38:10869 | ✓ 958ms | 否 | ✓ 663ms | ✓ 1587ms | ✓ 1210ms | http |
| 165.225.72.38:10514 | ✓ 867ms | 否 | ✓ 647ms | ✓ 1613ms | ✓ 1194ms | http |
| 165.225.72.38:11093 | ✓ 802ms | 否 | ✓ 631ms | ✓ 1620ms | ✓ 1196ms | http |
| 165.225.72.38:11526 | ✓ 765ms | 否 | ✓ 628ms | ✓ 1600ms | ✓ 1198ms | http |
| 165.225.72.38:11670 | ✓ 811ms | 否 | ✓ 628ms | ✓ 1548ms | ✓ 1208ms | http |
| 165.225.72.38:11023 | ✓ 675ms | ✓ 1940ms | ✓ 528ms | ✓ 1508ms | ✓ 1194ms | http |
| 165.225.72.38:10419 | ✓ 721ms | ✓ 1867ms | ✓ 531ms | ✓ 1507ms | ✓ 1191ms | http |
| 165.225.72.38:10456 | ✓ 604ms | ✓ 1806ms | ✓ 530ms | ✓ 1510ms | ✓ 1192ms | http |
| 38.34.179.35:8448 | ✓ 809ms | 否 | ✓ 729ms | ✓ 1086ms | 否 | http |
| 38.34.178.153:8448 | ✓ 524ms | ✓ 927ms | 否 | ✓ 1213ms | ✓ 1424ms | http |
| 45.136.130.177:8448 | ✓ 932ms | ✓ 894ms | ✓ 1639ms | ✓ 1693ms | ✓ 835ms | http |
| 101.47.73.135:3128 | ✓ 1465ms | 否 | 否 | ✓ 1631ms | ✓ 1536ms | http |
| 106.14.203.63:3333 | ✓ 826ms | ✓ 1924ms | ✓ 1155ms | 否 | ✓ 1392ms | http |
| 14.225.212.37:7890 | ✓ 1677ms | 否 | 否 | ✓ 1173ms | ✓ 849ms | http |
| 103.82.23.118:5247 | ✓ 1321ms | 否 | ✓ 1715ms | ✓ 1748ms | ✓ 1474ms | http |
| 113.176.92.71:3128 | ✓ 1825ms | ✓ 1321ms | ✓ 1284ms | ✓ 1251ms | ✓ 968ms | http |
| 45.149.92.147:5001 | ✓ 650ms | 否 | ✓ 809ms | 否 | ✓ 668ms | http |
| 103.217.216.71:1111 | ✓ 1783ms | 否 | 否 | ✓ 1909ms | ✓ 1444ms | http |
| 103.113.70.189:1081 | ✓ 876ms | ✓ 1064ms | 否 | ✓ 1140ms | ✓ 827ms | http |

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
