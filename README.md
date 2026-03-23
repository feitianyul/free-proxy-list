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

最后更新：2026-03-23 09:52:25 UTC（2026-03-23 17:52:25 UTC+8）

**代理总数：86**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 43.99.54.236:5555 | ✓ 640ms | ✓ 1556ms | ✓ 642ms | ✓ 812ms | ✓ 647ms | http |
| 147.161.210.140:8800 | ✓ 1566ms | ✓ 1380ms | ✓ 686ms | ✓ 901ms | ✓ 972ms | http |
| 1.231.81.166:3128 | ✓ 1596ms | ✓ 1002ms | 否 | ✓ 955ms | ✓ 1281ms | http |
| 219.117.204.211:7799 | ✓ 1566ms | 否 | ✓ 1011ms | ✓ 1168ms | ✓ 1492ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1353ms | ✓ 1043ms | ✓ 1273ms | ✓ 1320ms | http |
| 101.47.73.135:3128 | ✓ 1064ms | 否 | ✓ 1232ms | 否 | ✓ 1524ms | http |
| 45.167.124.52:8080 | ✓ 849ms | ✓ 1858ms | 否 | 否 | ✓ 1498ms | http |
| 104.129.202.127:10810 | ✓ 492ms | ✓ 770ms | ✓ 845ms | 否 | 否 | http |
| 104.129.202.127:12354 | ✓ 454ms | ✓ 747ms | ✓ 868ms | 否 | 否 | http |
| 35.225.22.61:80 | ✓ 689ms | ✓ 1593ms | 否 | 否 | ✓ 992ms | http |
| 183.249.5.117:22222 | ✓ 1304ms | 否 | ✓ 1608ms | ✓ 1864ms | ✓ 1501ms | http |
| 167.103.34.108:8800 | ✓ 1153ms | 否 | ✓ 1160ms | ✓ 1375ms | ✓ 1251ms | http |
| 167.103.31.122:8800 | ✓ 1892ms | 否 | ✓ 1795ms | 否 | ✓ 1973ms | http |
| 217.76.245.80:999 | ✓ 927ms | 否 | ✓ 1207ms | ✓ 1636ms | ✓ 1457ms | http |
| 83.219.250.8:62920 | ✓ 762ms | ✓ 1963ms | ✓ 1114ms | 否 | 否 | http |
| 137.220.151.110:6005 | ✓ 1447ms | ✓ 1909ms | ✓ 739ms | ✓ 1102ms | ✓ 885ms | http |
| 120.92.212.16:8890 | ✓ 1055ms | 否 | ✓ 1254ms | 否 | ✓ 1761ms | http |
| 147.161.239.240:8800 | ✓ 904ms | ✓ 1974ms | ✓ 1619ms | ✓ 1850ms | ✓ 1223ms | http |
| 45.136.198.40:3128 | ✓ 864ms | 否 | ✓ 1660ms | 否 | ✓ 1851ms | http |
| 101.43.127.100:8877 | ✓ 1652ms | ✓ 1595ms | ✓ 1473ms | 否 | ✓ 1852ms | http |
| 218.89.134.230:3333 | 否 | ✓ 1623ms | 否 | ✓ 1679ms | ✓ 1301ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 708ms | ✓ 1532ms | ✓ 1448ms | http |
| 34.101.184.164:3128 | ✓ 1605ms | 否 | ✓ 1476ms | ✓ 1642ms | ✓ 954ms | http |
| 194.67.99.223:1080 | ✓ 755ms | 否 | ✓ 1688ms | ✓ 1904ms | ✓ 1444ms | http |
| 47.101.149.27:9010 | ✓ 1276ms | ✓ 1419ms | 否 | ✓ 1976ms | 否 | http |
| 103.183.10.169:3125 | ✓ 1975ms | 否 | 否 | ✓ 1476ms | ✓ 1472ms | http |
| 45.144.28.81:10808 | ✓ 597ms | ✓ 1557ms | ✓ 1317ms | ✓ 1978ms | ✓ 1593ms | http |
| 104.168.158.236:10808 | ✓ 933ms | ✓ 1771ms | 否 | 否 | ✓ 1223ms | http |
| 59.46.216.131:30001 | ✓ 1044ms | ✓ 1371ms | 否 | ✓ 1356ms | 否 | http |
| 20.120.225.109:3128 | ✓ 434ms | ✓ 1783ms | ✓ 592ms | ✓ 1185ms | ✓ 754ms | http |
| 38.145.220.198:8448 | ✓ 714ms | ✓ 1887ms | ✓ 578ms | ✓ 693ms | ✓ 733ms | http |
| 45.136.130.187:8452 | ✓ 1887ms | 否 | ✓ 849ms | ✓ 668ms | ✓ 673ms | http |
| 38.34.179.70:8450 | ✓ 1907ms | 否 | ✓ 1814ms | ✓ 774ms | ✓ 848ms | http |
| 120.92.212.16:7890 | ✓ 1404ms | 否 | ✓ 969ms | ✓ 1250ms | 否 | http |
| 222.228.171.92:8080 | ✓ 1367ms | 否 | 否 | ✓ 1241ms | ✓ 989ms | http |
| 45.88.0.116:3128 | ✓ 1732ms | 否 | ✓ 827ms | 否 | ✓ 1728ms | http |
| 45.88.0.115:3128 | ✓ 593ms | 否 | ✓ 582ms | ✓ 1572ms | ✓ 1156ms | http |
| 222.184.48.251:22222 | ✓ 1264ms | ✓ 1821ms | ✓ 1054ms | ✓ 1606ms | 否 | http |
| 101.32.244.83:8080 | ✓ 972ms | 否 | ✓ 941ms | ✓ 1341ms | ✓ 1235ms | http |
| 121.43.196.213:8222 | ✓ 940ms | ✓ 1074ms | ✓ 829ms | ✓ 1111ms | ✓ 890ms | http |
| 121.43.196.210:8222 | ✓ 1013ms | ✓ 1041ms | ✓ 839ms | ✓ 1054ms | ✓ 882ms | http |
| 114.55.226.123:10086 | ✓ 1054ms | ✓ 1376ms | ✓ 998ms | ✓ 1244ms | ✓ 1020ms | http |
| 47.97.24.122:8222 | ✓ 936ms | ✓ 1237ms | ✓ 929ms | ✓ 1110ms | 否 | http |
| 103.139.138.194:3128 | ✓ 1904ms | 否 | ✓ 1680ms | ✓ 1889ms | ✓ 1400ms | http |
| 142.171.224.229:7890 | ✓ 689ms | ✓ 714ms | ✓ 826ms | ✓ 712ms | ✓ 518ms | http |
| 172.212.68.37:3128 | ✓ 466ms | 否 | ✓ 898ms | ✓ 1541ms | ✓ 1248ms | http |
| 58.220.95.8:10174 | ✓ 929ms | ✓ 1794ms | 否 | ✓ 1229ms | ✓ 923ms | http |
| 160.250.4.245:1 | ✓ 1859ms | 否 | 否 | ✓ 1446ms | ✓ 1143ms | http |
| 38.34.179.162:8451 | ✓ 870ms | 否 | ✓ 1249ms | ✓ 1038ms | ✓ 1175ms | http |
| 38.34.179.228:8449 | ✓ 828ms | ✓ 1962ms | ✓ 465ms | ✓ 703ms | ✓ 784ms | http |
| 183.98.143.134:8046 | ✓ 1022ms | ✓ 1305ms | ✓ 1602ms | ✓ 1246ms | ✓ 1662ms | http |
| 185.191.236.162:3128 | ✓ 803ms | 否 | ✓ 1078ms | ✓ 1868ms | ✓ 1275ms | http |
| 38.34.179.99:8444 | ✓ 217ms | ✓ 883ms | 否 | ✓ 1531ms | 否 | http |
| 38.34.179.54:8446 | ✓ 647ms | ✓ 686ms | ✓ 1824ms | ✓ 1821ms | ✓ 734ms | http |
| 166.88.55.83:7890 | ✓ 628ms | ✓ 1072ms | ✓ 634ms | ✓ 781ms | ✓ 637ms | http |
| 183.249.5.105:22222 | ✓ 883ms | ✓ 1089ms | ✓ 725ms | ✓ 947ms | ✓ 732ms | http |
| 202.141.161.53:30001 | ✓ 1456ms | ✓ 1950ms | ✓ 1713ms | ✓ 1217ms | ✓ 1097ms | http |
| 106.117.208.101:7890 | ✓ 908ms | ✓ 1273ms | ✓ 1084ms | ✓ 1262ms | ✓ 1614ms | http |
| 137.220.150.170:6005 | ✓ 1386ms | 否 | ✓ 1442ms | ✓ 1337ms | 否 | http |
| 47.77.193.180:1080 | ✓ 149ms | ✓ 1289ms | ✓ 231ms | ✓ 709ms | ✓ 564ms | http |
| 45.140.147.155:1082 | ✓ 652ms | ✓ 1909ms | ✓ 1152ms | ✓ 1600ms | ✓ 1261ms | http |
| 207.254.71.62:8088 | ✓ 1348ms | 否 | ✓ 1575ms | 否 | ✓ 1913ms | http |
| 43.165.195.107:3128 | ✓ 817ms | 否 | ✓ 876ms | ✓ 1179ms | ✓ 934ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1771ms | ✓ 1649ms | ✓ 1436ms | http |
| 183.249.5.110:22222 | ✓ 827ms | ✓ 1299ms | ✓ 1127ms | ✓ 1140ms | ✓ 748ms | http |
| 103.84.95.54:7890 | ✓ 1108ms | 否 | 否 | ✓ 1636ms | ✓ 670ms | http |
| 222.184.48.252:22222 | ✓ 895ms | ✓ 1198ms | ✓ 1885ms | 否 | ✓ 1756ms | http |
| 106.75.15.167:7890 | ✓ 1526ms | ✓ 1674ms | 否 | ✓ 1366ms | 否 | http |
| 181.78.44.63:999 | 否 | 否 | ✓ 1311ms | ✓ 1557ms | ✓ 1374ms | http |
| 137.220.150.104:6005 | 否 | 否 | ✓ 1038ms | ✓ 1122ms | ✓ 1090ms | http |
| 211.95.152.50:45046 | ✓ 1731ms | ✓ 1459ms | ✓ 1772ms | 否 | 否 | http |
| 38.145.220.11:8445 | ✓ 832ms | ✓ 1545ms | 否 | ✓ 1138ms | ✓ 716ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1330ms | ✓ 1836ms | ✓ 1287ms | http |
| 45.149.92.147:5001 | ✓ 638ms | 否 | ✓ 635ms | ✓ 796ms | ✓ 699ms | http |
| 38.145.220.37:8453 | ✓ 1599ms | ✓ 1851ms | ✓ 228ms | ✓ 1146ms | 否 | http |
| 38.145.220.14:8453 | ✓ 1601ms | ✓ 1828ms | ✓ 122ms | ✓ 1004ms | ✓ 1965ms | http |
| 38.145.220.36:8453 | ✓ 1600ms | 否 | ✓ 456ms | ✓ 1920ms | ✓ 1441ms | http |
| 14.145.134.227:10808 | ✓ 1662ms | ✓ 1209ms | ✓ 1044ms | ✓ 1165ms | ✓ 980ms | http |
| 121.230.8.208:1080 | ✓ 1781ms | ✓ 1392ms | ✓ 1828ms | 否 | 否 | http |
| 103.67.46.225:3125 | ✓ 1835ms | 否 | 否 | ✓ 1674ms | ✓ 1597ms | http |
| 38.145.203.35:8448 | ✓ 1712ms | ✓ 1451ms | ✓ 322ms | ✓ 1385ms | 否 | http |
| 45.88.0.117:3128 | 否 | 否 | ✓ 1414ms | ✓ 1859ms | ✓ 1516ms | http |
| 38.145.208.166:8451 | ✓ 435ms | ✓ 1178ms | ✓ 133ms | 否 | 否 | http |
| 137.220.150.22:6005 | ✓ 1578ms | 否 | ✓ 1371ms | 否 | ✓ 1827ms | http |
| 112.65.132.182:3128 | ✓ 847ms | 否 | ✓ 833ms | ✓ 1039ms | ✓ 816ms | http |
| 154.201.64.99:7890 | ✓ 1455ms | ✓ 1231ms | 否 | 否 | ✓ 1700ms | http |

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
