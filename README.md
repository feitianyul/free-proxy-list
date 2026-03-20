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

最后更新：2026-03-20 20:36:31 UTC（2026-03-21 04:36:31 UTC+8）

**代理总数：79**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 78 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 79 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 658ms | ✓ 1349ms | ✓ 884ms | ✓ 1023ms | ✓ 1859ms | http |
| 178.156.187.185:10001 | ✓ 365ms | ✓ 1402ms | ✓ 1206ms | 否 | ✓ 1489ms | http |
| 137.220.151.110:6005 | 否 | 否 | ✓ 1196ms | ✓ 1362ms | ✓ 1217ms | http |
| 174.138.24.77:1080 | ✓ 1920ms | 否 | ✓ 988ms | ✓ 1064ms | ✓ 1016ms | http |
| 113.160.132.26:8080 | ✓ 1653ms | ✓ 1374ms | ✓ 1491ms | ✓ 1401ms | ✓ 1355ms | http |
| 137.220.150.104:6005 | ✓ 1386ms | 否 | ✓ 1238ms | ✓ 1632ms | ✓ 1470ms | http |
| 137.220.150.170:6005 | ✓ 1237ms | ✓ 1877ms | ✓ 1809ms | ✓ 1943ms | ✓ 907ms | http |
| 167.103.34.108:8800 | ✓ 1536ms | 否 | ✓ 1536ms | ✓ 1568ms | ✓ 1612ms | http |
| 45.167.124.52:8080 | ✓ 1611ms | ✓ 1929ms | ✓ 1564ms | 否 | ✓ 1768ms | http |
| 35.225.22.61:80 | ✓ 403ms | 否 | ✓ 959ms | ✓ 1098ms | ✓ 1018ms | http |
| 38.34.179.27:8451 | ✓ 1775ms | ✓ 865ms | ✓ 320ms | ✓ 1024ms | ✓ 1695ms | http |
| 219.117.204.211:7799 | ✓ 1411ms | 否 | ✓ 512ms | 否 | ✓ 672ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1209ms | ✓ 1023ms | ✓ 1249ms | ✓ 997ms | http |
| 137.220.150.22:6005 | ✓ 1402ms | 否 | ✓ 958ms | ✓ 1910ms | ✓ 1753ms | http |
| 103.84.95.54:7890 | ✓ 1858ms | 否 | ✓ 1941ms | ✓ 1018ms | ✓ 1234ms | http |
| 137.220.150.152:6005 | ✓ 1987ms | 否 | 否 | ✓ 1243ms | ✓ 1187ms | http |
| 45.140.147.155:1081 | ✓ 794ms | 否 | ✓ 1158ms | ✓ 1881ms | 否 | http |
| 222.184.48.251:22222 | ✓ 1035ms | ✓ 1178ms | ✓ 1542ms | 否 | 否 | http |
| 162.240.154.26:3128 | ✓ 847ms | ✓ 1604ms | ✓ 843ms | 否 | ✓ 982ms | http |
| 133.242.138.34:8100 | 否 | 否 | ✓ 1330ms | ✓ 1337ms | ✓ 1087ms | http |
| 116.80.65.80:3172 | ✓ 1533ms | 否 | ✓ 1521ms | ✓ 1871ms | 否 | http |
| 183.249.5.117:22222 | ✓ 769ms | ✓ 1167ms | ✓ 760ms | 否 | 否 | http |
| 139.159.99.242:8080 | ✓ 807ms | ✓ 1031ms | 否 | ✓ 1060ms | ✓ 816ms | http |
| 120.92.212.16:8890 | ✓ 960ms | ✓ 1230ms | 否 | ✓ 1505ms | 否 | http |
| 45.136.131.38:8449 | ✓ 811ms | ✓ 1683ms | ✓ 394ms | ✓ 906ms | ✓ 1244ms | http |
| 101.43.127.100:8877 | ✓ 941ms | ✓ 1122ms | ✓ 949ms | ✓ 1067ms | ✓ 829ms | http |
| 147.161.239.240:8800 | ✓ 1250ms | ✓ 1700ms | ✓ 997ms | ✓ 1690ms | ✓ 1611ms | http |
| 91.238.105.64:2024 | ✓ 1275ms | ✓ 1673ms | ✓ 1661ms | 否 | ✓ 1983ms | http |
| 121.230.9.19:1080 | ✓ 1127ms | ✓ 1522ms | ✓ 1236ms | ✓ 1524ms | 否 | http |
| 38.145.208.182:8451 | ✓ 266ms | ✓ 1285ms | ✓ 543ms | ✓ 922ms | ✓ 605ms | http |
| 38.34.179.58:8447 | ✓ 714ms | 否 | ✓ 504ms | ✓ 1051ms | ✓ 888ms | http |
| 167.103.31.122:8800 | ✓ 1371ms | 否 | ✓ 1393ms | ✓ 1669ms | ✓ 1678ms | http |
| 116.80.49.165:3172 | 否 | 否 | ✓ 1501ms | ✓ 1848ms | ✓ 1675ms | http |
| 192.71.213.85:5555 | ✓ 1591ms | 否 | ✓ 1331ms | ✓ 1824ms | 否 | http |
| 147.45.67.148:8080 | 否 | 否 | ✓ 1196ms | ✓ 1788ms | ✓ 1409ms | http |
| 14.225.212.37:7890 | ✓ 1826ms | 否 | ✓ 846ms | ✓ 1114ms | 否 | http |
| 45.93.30.241:6005 | ✓ 1466ms | 否 | 否 | ✓ 1830ms | ✓ 1674ms | http |
| 1.231.81.166:3128 | ✓ 1939ms | 否 | ✓ 1991ms | 否 | ✓ 1950ms | http |
| 115.231.181.40:8128 | ✓ 910ms | 否 | ✓ 853ms | ✓ 1293ms | 否 | http |
| 38.34.179.97:8448 | ✓ 1507ms | ✓ 686ms | ✓ 580ms | ✓ 1973ms | ✓ 1245ms | http |
| 59.46.216.131:30001 | ✓ 991ms | 否 | ✓ 1086ms | ✓ 1509ms | 否 | http |
| 88.80.150.82:8080 | ✓ 1980ms | 否 | ✓ 1209ms | 否 | ✓ 1880ms | https |
| 43.99.54.236:5555 | ✓ 1816ms | ✓ 958ms | ✓ 712ms | ✓ 837ms | ✓ 681ms | http |
| 116.80.49.162:3172 | 否 | 否 | ✓ 1835ms | ✓ 1946ms | ✓ 1845ms | http |
| 34.96.238.40:8080 | 否 | 否 | ✓ 922ms | ✓ 1044ms | ✓ 1372ms | http |
| 38.34.179.20:8445 | ✓ 524ms | ✓ 1227ms | ✓ 1757ms | ✓ 1378ms | ✓ 716ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1917ms | ✓ 972ms | ✓ 1145ms | ✓ 843ms | http |
| 114.237.77.231:1080 | ✓ 963ms | ✓ 1138ms | ✓ 982ms | ✓ 1290ms | 否 | http |
| 106.117.208.101:7890 | ✓ 1080ms | ✓ 1294ms | ✓ 1263ms | ✓ 1419ms | ✓ 1041ms | http |
| 137.184.14.135:3128 | ✓ 1370ms | ✓ 830ms | ✓ 924ms | ✓ 950ms | 否 | http |
| 45.136.198.40:3128 | ✓ 1232ms | 否 | ✓ 1083ms | 否 | ✓ 1566ms | http |
| 51.77.210.150:3128 | ✓ 1603ms | 否 | ✓ 1778ms | 否 | ✓ 1848ms | http |
| 150.249.255.91:3128 | ✓ 1444ms | ✓ 953ms | ✓ 623ms | ✓ 917ms | ✓ 736ms | http |
| 38.34.179.96:8451 | ✓ 955ms | ✓ 685ms | ✓ 654ms | ✓ 1185ms | ✓ 1904ms | http |
| 106.75.15.167:7890 | 否 | 否 | ✓ 1585ms | ✓ 1574ms | ✓ 1193ms | http |
| 58.220.95.8:10174 | 否 | 否 | ✓ 1040ms | ✓ 1285ms | ✓ 960ms | http |
| 183.249.5.110:22222 | ✓ 847ms | ✓ 1035ms | ✓ 1842ms | ✓ 996ms | ✓ 850ms | http |
| 210.223.44.230:3128 | ✓ 1864ms | 否 | ✓ 1873ms | ✓ 1144ms | ✓ 942ms | http |
| 194.67.99.223:1080 | ✓ 779ms | 否 | ✓ 1749ms | 否 | ✓ 1645ms | http |
| 34.101.184.164:3128 | ✓ 1723ms | 否 | ✓ 1103ms | ✓ 1453ms | ✓ 1627ms | http |
| 38.55.104.182:6005 | ✓ 1103ms | 否 | ✓ 1580ms | ✓ 1107ms | ✓ 902ms | http |
| 38.34.183.225:8450 | ✓ 1874ms | ✓ 1185ms | 否 | ✓ 1952ms | 否 | http |
| 38.34.179.16:8451 | 否 | ✓ 1643ms | 否 | ✓ 1042ms | ✓ 950ms | http |
| 38.34.179.30:8451 | 否 | ✓ 1622ms | 否 | ✓ 1050ms | ✓ 964ms | http |
| 45.93.28.159:6005 | ✓ 1141ms | 否 | 否 | ✓ 1855ms | ✓ 1157ms | http |
| 121.232.73.214:1080 | ✓ 1035ms | ✓ 1208ms | ✓ 900ms | ✓ 1233ms | ✓ 1006ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1754ms | ✓ 1492ms | ✓ 1324ms | http |
| 38.34.179.105:8449 | ✓ 757ms | ✓ 1157ms | ✓ 1982ms | ✓ 753ms | ✓ 1170ms | http |
| 116.80.65.76:3172 | ✓ 1603ms | 否 | 否 | ✓ 1950ms | ✓ 1811ms | http |
| 183.249.5.105:22222 | ✓ 832ms | ✓ 1543ms | ✓ 720ms | ✓ 1034ms | ✓ 1521ms | http |
| 47.95.231.180:8084 | ✓ 929ms | ✓ 1222ms | ✓ 927ms | ✓ 1230ms | ✓ 962ms | http |
| 106.14.203.63:3333 | ✓ 857ms | ✓ 1468ms | 否 | 否 | ✓ 1719ms | http |
| 45.186.6.104:3128 | ✓ 1851ms | ✓ 1698ms | ✓ 1710ms | 否 | 否 | http |
| 45.136.131.59:8450 | ✓ 333ms | ✓ 752ms | ✓ 1041ms | ✓ 731ms | ✓ 575ms | http |
| 59.8.203.55:80 | ✓ 1636ms | ✓ 1218ms | ✓ 1249ms | ✓ 999ms | ✓ 818ms | http |
| 172.212.68.37:3128 | ✓ 360ms | 否 | ✓ 703ms | ✓ 1820ms | ✓ 1135ms | http |
| 183.249.5.111:22222 | ✓ 1754ms | ✓ 1046ms | ✓ 721ms | ✓ 955ms | 否 | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 822ms | ✓ 1235ms | ✓ 920ms | http |
| 38.34.179.60:8450 | ✓ 958ms | ✓ 1683ms | ✓ 1465ms | ✓ 885ms | 否 | http |

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
