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

最后更新：2026-03-24 14:12:01 UTC（2026-03-24 22:12:01 UTC+8）

**代理总数：84**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 84 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 84 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 1043ms | 否 | ✓ 1090ms | 否 | ✓ 1045ms | http |
| 8.217.106.71:8888 | ✓ 660ms | 否 | ✓ 659ms | ✓ 846ms | ✓ 668ms | http |
| 147.161.210.140:8800 | ✓ 1624ms | 否 | ✓ 848ms | ✓ 1809ms | ✓ 1268ms | http |
| 137.220.151.110:6005 | ✓ 1758ms | 否 | ✓ 847ms | ✓ 1257ms | ✓ 1184ms | http |
| 167.103.115.102:8800 | ✓ 1696ms | 否 | ✓ 1044ms | ✓ 1135ms | ✓ 1243ms | http |
| 202.141.161.53:10808 | ✓ 1034ms | 否 | ✓ 1264ms | 否 | ✓ 1804ms | http |
| 167.103.34.108:8800 | ✓ 1748ms | 否 | ✓ 1717ms | ✓ 1413ms | ✓ 1528ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 1196ms | ✓ 1420ms | ✓ 1104ms | http |
| 47.77.193.180:1080 | ✓ 193ms | 否 | ✓ 475ms | ✓ 750ms | ✓ 557ms | http |
| 166.88.55.83:7890 | ✓ 648ms | 否 | ✓ 676ms | ✓ 834ms | ✓ 665ms | http |
| 142.171.224.229:7890 | ✓ 1066ms | 否 | ✓ 1909ms | ✓ 930ms | ✓ 1694ms | http |
| 103.84.95.54:7890 | ✓ 1168ms | 否 | ✓ 1938ms | ✓ 1327ms | ✓ 1120ms | http |
| 155.212.132.241:3128 | ✓ 722ms | 否 | ✓ 1759ms | 否 | ✓ 1436ms | http |
| 1.231.81.166:3128 | ✓ 1195ms | 否 | ✓ 1759ms | ✓ 1218ms | ✓ 1356ms | http |
| 167.103.31.122:8800 | ✓ 1639ms | 否 | ✓ 1765ms | 否 | ✓ 1974ms | http |
| 38.145.208.213:8450 | ✓ 1488ms | 否 | ✓ 1658ms | ✓ 1545ms | ✓ 1790ms | http |
| 38.145.208.206:8448 | ✓ 1511ms | 否 | ✓ 1719ms | ✓ 1523ms | ✓ 1759ms | http |
| 45.167.124.52:8080 | ✓ 692ms | 否 | ✓ 833ms | ✓ 1667ms | ✓ 1395ms | http |
| 147.161.239.240:8800 | ✓ 673ms | 否 | ✓ 1205ms | ✓ 1788ms | ✓ 1631ms | http |
| 181.41.201.85:3128 | ✓ 1377ms | 否 | ✓ 961ms | 否 | ✓ 1626ms | http |
| 120.92.212.16:8890 | ✓ 1234ms | 否 | ✓ 1895ms | ✓ 1264ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1214ms | ✓ 1457ms | ✓ 953ms | ✓ 1943ms | ✓ 1002ms | http |
| 160.250.5.22:1 | ✓ 1588ms | 否 | 否 | ✓ 1285ms | ✓ 1080ms | http |
| 101.43.127.100:8877 | ✓ 879ms | ✓ 1438ms | ✓ 1633ms | ✓ 1187ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1320ms | ✓ 1363ms | 否 | ✓ 1682ms | ✓ 1004ms | http |
| 58.220.95.8:10174 | 否 | 否 | ✓ 1960ms | ✓ 1409ms | ✓ 973ms | http |
| 218.89.134.230:3333 | ✓ 1595ms | 否 | 否 | ✓ 1661ms | ✓ 1342ms | http |
| 137.220.150.104:6005 | ✓ 821ms | 否 | ✓ 802ms | ✓ 1269ms | 否 | http |
| 217.77.102.18:3128 | ✓ 1263ms | 否 | ✓ 1552ms | ✓ 1945ms | ✓ 1711ms | http |
| 20.78.26.206:8561 | 否 | ✓ 1428ms | ✓ 502ms | ✓ 803ms | ✓ 629ms | http |
| 20.27.14.220:8561 | 否 | ✓ 1772ms | ✓ 680ms | ✓ 832ms | ✓ 630ms | http |
| 20.78.118.91:8561 | 否 | 否 | ✓ 487ms | ✓ 797ms | ✓ 628ms | http |
| 20.210.76.104:8561 | 否 | 否 | ✓ 490ms | ✓ 804ms | ✓ 630ms | http |
| 47.101.159.19:8899 | ✓ 882ms | ✓ 1102ms | ✓ 1265ms | 否 | ✓ 891ms | http |
| 62.113.119.14:8080 | ✓ 788ms | ✓ 1816ms | ✓ 1083ms | ✓ 1651ms | ✓ 1336ms | http |
| 104.243.46.122:3128 | 否 | 否 | ✓ 1197ms | ✓ 1594ms | ✓ 1272ms | http |
| 160.250.4.245:1 | ✓ 1511ms | 否 | ✓ 1250ms | ✓ 1294ms | ✓ 996ms | http |
| 222.228.171.92:8080 | 否 | 否 | ✓ 978ms | ✓ 1232ms | ✓ 1527ms | http |
| 45.136.131.34:8449 | ✓ 971ms | ✓ 962ms | ✓ 354ms | ✓ 790ms | ✓ 1299ms | http |
| 85.208.51.165:443 | 否 | 否 | ✓ 1662ms | ✓ 1787ms | ✓ 1354ms | http |
| 83.219.250.8:62920 | ✓ 850ms | ✓ 1689ms | ✓ 1615ms | 否 | ✓ 1920ms | http |
| 207.254.71.62:8088 | ✓ 1137ms | 否 | 否 | ✓ 1861ms | ✓ 1718ms | http |
| 103.184.99.194:8080 | ✓ 1868ms | 否 | 否 | ✓ 1422ms | ✓ 1440ms | http |
| 144.31.79.117:8888 | ✓ 818ms | 否 | ✓ 1390ms | 否 | ✓ 1500ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1413ms | ✓ 1400ms | ✓ 1053ms | http |
| 193.233.22.29:10808 | ✓ 901ms | 否 | ✓ 1216ms | ✓ 1846ms | 否 | http |
| 20.27.15.111:8561 | ✓ 508ms | ✓ 1080ms | ✓ 668ms | ✓ 859ms | ✓ 628ms | http |
| 38.34.183.13:8445 | ✓ 158ms | ✓ 846ms | ✓ 746ms | ✓ 974ms | ✓ 1501ms | http |
| 184.72.55.45:3128 | ✓ 387ms | ✓ 1039ms | ✓ 1173ms | ✓ 1029ms | ✓ 1016ms | http |
| 45.136.130.187:8452 | ✓ 206ms | ✓ 1026ms | ✓ 314ms | ✓ 1329ms | ✓ 1480ms | http |
| 20.27.13.35:8561 | 否 | ✓ 894ms | ✓ 489ms | ✓ 859ms | ✓ 628ms | http |
| 38.145.203.97:8445 | ✓ 621ms | ✓ 1918ms | ✓ 161ms | ✓ 774ms | ✓ 588ms | http |
| 38.34.179.162:8453 | ✓ 313ms | ✓ 1238ms | ✓ 1604ms | ✓ 761ms | ✓ 960ms | http |
| 38.34.179.22:8445 | ✓ 390ms | ✓ 1142ms | ✓ 1432ms | ✓ 1334ms | ✓ 671ms | http |
| 38.145.203.105:8447 | ✓ 1170ms | ✓ 1764ms | ✓ 272ms | ✓ 1079ms | ✓ 1951ms | http |
| 38.145.220.6:8446 | ✓ 675ms | ✓ 1353ms | ✓ 340ms | ✓ 1070ms | ✓ 1939ms | http |
| 101.47.73.135:3128 | ✓ 1913ms | 否 | ✓ 1718ms | ✓ 1369ms | 否 | http |
| 57.128.188.167:9163 | ✓ 1727ms | 否 | ✓ 1705ms | 否 | ✓ 1745ms | http |
| 117.1.136.57:6620 | ✓ 1498ms | 否 | ✓ 1486ms | ✓ 1686ms | ✓ 1396ms | http |
| 137.220.150.22:6005 | ✓ 891ms | 否 | ✓ 1411ms | ✓ 1253ms | ✓ 1335ms | http |
| 150.241.71.15:1080 | ✓ 1236ms | ✓ 1731ms | ✓ 1949ms | 否 | 否 | http |
| 8.219.97.248:80 | ✓ 1344ms | 否 | ✓ 1364ms | ✓ 1920ms | 否 | http |
| 202.141.161.53:30001 | ✓ 1979ms | 否 | 否 | ✓ 1989ms | ✓ 1081ms | http |
| 20.27.15.49:8561 | ✓ 509ms | ✓ 1051ms | ✓ 688ms | ✓ 870ms | ✓ 652ms | http |
| 20.210.76.175:8561 | ✓ 508ms | ✓ 1481ms | ✓ 508ms | ✓ 849ms | ✓ 640ms | http |
| 20.210.76.178:8561 | ✓ 510ms | ✓ 1522ms | ✓ 523ms | ✓ 842ms | ✓ 649ms | http |
| 45.136.130.189:8450 | ✓ 821ms | ✓ 1907ms | 否 | ✓ 1260ms | ✓ 877ms | http |
| 106.75.15.167:7890 | ✓ 1020ms | 否 | ✓ 1106ms | ✓ 1231ms | 否 | http |
| 223.16.170.103:80 | ✓ 1228ms | 否 | ✓ 1133ms | ✓ 1122ms | 否 | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1044ms | ✓ 1412ms | ✓ 1024ms | http |
| 137.220.150.170:6005 | ✓ 1228ms | 否 | ✓ 792ms | ✓ 1184ms | ✓ 938ms | http |
| 45.129.141.143:3128 | ✓ 1238ms | 否 | ✓ 1751ms | 否 | ✓ 1891ms | http |
| 20.210.39.153:8561 | ✓ 1398ms | ✓ 1281ms | ✓ 513ms | ✓ 796ms | ✓ 623ms | http |
| 20.27.11.248:8561 | ✓ 1392ms | ✓ 1513ms | ✓ 485ms | ✓ 800ms | ✓ 692ms | http |
| 62.171.161.88:2018 | ✓ 1124ms | ✓ 1830ms | ✓ 1700ms | ✓ 1821ms | ✓ 1541ms | http |
| 103.69.84.106:3131 | ✓ 1816ms | 否 | ✓ 1204ms | ✓ 1237ms | ✓ 990ms | http |
| 115.231.181.40:8128 | ✓ 1178ms | ✓ 1672ms | 否 | ✓ 1643ms | 否 | http |
| 61.52.131.172:8443 | ✓ 846ms | ✓ 1166ms | ✓ 928ms | ✓ 1169ms | ✓ 929ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1436ms | ✓ 1275ms | ✓ 1253ms | ✓ 978ms | http |
| 45.136.198.40:3128 | ✓ 1361ms | ✓ 1734ms | ✓ 1877ms | 否 | ✓ 1898ms | http |
| 137.220.150.152:6005 | ✓ 1580ms | 否 | ✓ 1448ms | ✓ 1262ms | ✓ 1667ms | http |
| 144.31.52.77:3128 | ✓ 945ms | 否 | ✓ 1478ms | 否 | ✓ 1755ms | http |
| 137.184.6.37:3128 | ✓ 1380ms | ✓ 1155ms | ✓ 1061ms | ✓ 759ms | ✓ 1835ms | http |
| 103.183.10.169:3125 | 否 | 否 | ✓ 1314ms | ✓ 1506ms | ✓ 1968ms | http |

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
