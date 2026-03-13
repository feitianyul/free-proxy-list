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

最后更新：2026-03-13 05:43:02 UTC（2026-03-13 13:43:02 UTC+8）

**代理总数：90**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 89 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 90 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.130.175:8443 | 否 | ✓ 1167ms | ✓ 659ms | ✓ 672ms | ✓ 512ms | http |
| 45.136.131.47:8443 | 否 | ✓ 1164ms | ✓ 790ms | ✓ 653ms | ✓ 507ms | http |
| 46.183.25.8:443 | 否 | 否 | ✓ 1251ms | ✓ 864ms | ✓ 1205ms | http |
| 45.136.131.63:8443 | 否 | 否 | ✓ 665ms | ✓ 970ms | ✓ 517ms | http |
| 103.84.95.54:7890 | ✓ 659ms | 否 | ✓ 741ms | ✓ 1532ms | 否 | http |
| 1.231.81.166:3128 | 否 | 否 | ✓ 843ms | ✓ 844ms | ✓ 755ms | http |
| 113.160.132.26:8080 | ✓ 1635ms | ✓ 1364ms | ✓ 1338ms | ✓ 1158ms | ✓ 958ms | http |
| 152.42.213.210:8080 | ✓ 1965ms | 否 | ✓ 1379ms | ✓ 1051ms | ✓ 803ms | http |
| 205.209.118.30:3138 | ✓ 1100ms | 否 | ✓ 1276ms | 否 | ✓ 1066ms | http |
| 178.236.245.17:3128 | ✓ 1602ms | 否 | ✓ 1750ms | 否 | ✓ 1878ms | http |
| 185.191.236.162:3128 | ✓ 1934ms | ✓ 1864ms | 否 | 否 | ✓ 1905ms | http |
| 45.167.124.52:8080 | ✓ 1554ms | 否 | ✓ 1151ms | 否 | ✓ 1762ms | http |
| 120.92.212.16:7890 | ✓ 933ms | ✓ 1138ms | ✓ 947ms | ✓ 1178ms | ✓ 969ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1388ms | ✓ 1436ms | ✓ 1411ms | ✓ 1187ms | http |
| 217.76.245.80:999 | ✓ 1149ms | 否 | ✓ 1182ms | ✓ 1413ms | ✓ 1266ms | http |
| 67.169.98.211:443 | ✓ 1827ms | 否 | ✓ 1448ms | ✓ 1499ms | ✓ 787ms | http |
| 160.238.65.4:3128 | ✓ 1210ms | 否 | ✓ 1263ms | 否 | ✓ 1506ms | http |
| 160.238.65.3:3128 | ✓ 1210ms | ✓ 1851ms | ✓ 1397ms | 否 | ✓ 1521ms | http |
| 160.238.65.5:3128 | ✓ 1210ms | 否 | ✓ 1250ms | 否 | ✓ 1525ms | http |
| 160.238.65.8:3128 | ✓ 1209ms | ✓ 1673ms | ✓ 1577ms | 否 | ✓ 1516ms | http |
| 160.238.65.2:3128 | ✓ 1208ms | ✓ 1811ms | ✓ 1440ms | 否 | ✓ 1514ms | http |
| 160.238.65.7:3128 | ✓ 1209ms | 否 | ✓ 1251ms | 否 | ✓ 1515ms | http |
| 160.238.65.9:3128 | ✓ 1208ms | 否 | ✓ 1250ms | 否 | ✓ 1535ms | http |
| 160.238.65.6:3128 | ✓ 1209ms | ✓ 1808ms | ✓ 1440ms | 否 | ✓ 1528ms | http |
| 116.80.47.62:3172 | 否 | 否 | ✓ 1560ms | ✓ 1816ms | ✓ 1599ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1391ms | ✓ 1292ms | ✓ 1088ms | http |
| 103.113.70.189:1081 | ✓ 403ms | ✓ 1174ms | 否 | ✓ 1359ms | 否 | http |
| 103.82.93.219:3128 | ✓ 1588ms | 否 | ✓ 1087ms | ✓ 1237ms | ✓ 993ms | http |
| 193.168.173.136:443 | ✓ 1193ms | 否 | ✓ 1613ms | ✓ 1965ms | 否 | http |
| 81.70.169.194:80 | ✓ 1732ms | ✓ 1578ms | ✓ 1032ms | ✓ 1168ms | ✓ 932ms | http |
| 101.43.255.96:80 | ✓ 928ms | ✓ 1435ms | ✓ 1081ms | 否 | 否 | http |
| 43.167.227.161:1080 | 否 | 否 | ✓ 806ms | ✓ 1938ms | ✓ 707ms | http |
| 190.9.109.198:999 | ✓ 1018ms | 否 | ✓ 1288ms | ✓ 1452ms | ✓ 1435ms | http |
| 45.136.130.223:8443 | ✓ 1708ms | ✓ 1181ms | ✓ 350ms | ✓ 686ms | ✓ 510ms | http |
| 45.136.130.188:8443 | ✓ 1675ms | ✓ 1255ms | ✓ 295ms | ✓ 653ms | ✓ 746ms | http |
| 8.219.97.248:80 | ✓ 1258ms | 否 | ✓ 1094ms | ✓ 1345ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1575ms | 否 | ✓ 1580ms | ✓ 1696ms | ✓ 996ms | http |
| 116.80.49.156:3172 | ✓ 1669ms | 否 | 否 | ✓ 1798ms | ✓ 1687ms | http |
| 171.251.172.78:5102 | 否 | 否 | ✓ 1697ms | ✓ 1412ms | ✓ 1913ms | http |
| 45.136.130.191:8443 | ✓ 81ms | ✓ 702ms | ✓ 79ms | ✓ 684ms | ✓ 503ms | http |
| 47.101.149.27:9010 | ✓ 1295ms | ✓ 1229ms | ✓ 1289ms | 否 | 否 | http |
| 39.104.201.40:7890 | ✓ 1170ms | ✓ 1228ms | ✓ 862ms | ✓ 1191ms | ✓ 915ms | http |
| 160.250.4.245:1 | ✓ 944ms | 否 | ✓ 1293ms | ✓ 1287ms | ✓ 1401ms | http |
| 115.231.181.40:8128 | ✓ 898ms | 否 | ✓ 1638ms | 否 | ✓ 861ms | http |
| 223.16.170.103:80 | ✓ 848ms | 否 | ✓ 944ms | ✓ 1005ms | ✓ 1032ms | http |
| 171.251.172.78:5107 | ✓ 1902ms | 否 | 否 | ✓ 1456ms | ✓ 1341ms | http |
| 116.80.65.77:3172 | ✓ 1627ms | 否 | 否 | ✓ 1778ms | ✓ 1670ms | http |
| 85.208.108.43:2094 | ✓ 622ms | 否 | 否 | ✓ 1309ms | ✓ 959ms | http |
| 35.225.22.61:80 | ✓ 886ms | 否 | 否 | ✓ 1340ms | ✓ 1103ms | http |
| 138.124.53.25:7443 | ✓ 1011ms | ✓ 1979ms | ✓ 1022ms | ✓ 1824ms | 否 | http |
| 116.80.49.166:3172 | ✓ 1447ms | 否 | ✓ 1455ms | ✓ 1783ms | 否 | http |
| 116.80.96.107:3172 | ✓ 1573ms | 否 | 否 | ✓ 1748ms | ✓ 1650ms | http |
| 116.80.49.163:3172 | ✓ 1837ms | 否 | ✓ 1473ms | ✓ 1777ms | ✓ 1639ms | http |
| 116.80.49.165:3172 | 否 | 否 | ✓ 1487ms | ✓ 1775ms | ✓ 1635ms | http |
| 116.80.49.159:3172 | ✓ 1446ms | 否 | ✓ 1457ms | ✓ 1803ms | 否 | http |
| 178.236.245.59:3128 | ✓ 1094ms | 否 | ✓ 863ms | ✓ 1854ms | ✓ 1454ms | http |
| 45.140.147.82:1081 | ✓ 1095ms | 否 | ✓ 1824ms | 否 | ✓ 1502ms | http |
| 220.197.44.36:3128 | ✓ 1680ms | 否 | ✓ 1354ms | 否 | ✓ 1690ms | http |
| 86.53.183.16:1080 | ✓ 950ms | 否 | ✓ 1404ms | 否 | ✓ 1640ms | http |
| 165.227.5.10:8888 | ✓ 1862ms | ✓ 1083ms | ✓ 1358ms | 否 | 否 | http |
| 162.248.165.72:1080 | ✓ 1070ms | 否 | ✓ 755ms | 否 | ✓ 1272ms | http |
| 168.235.110.63:3128 | ✓ 1210ms | ✓ 1485ms | ✓ 1527ms | 否 | 否 | http |
| 89.185.85.138:1080 | ✓ 610ms | 否 | ✓ 1534ms | ✓ 1919ms | ✓ 1004ms | http |
| 194.116.191.181:3128 | ✓ 941ms | 否 | ✓ 1498ms | ✓ 1890ms | ✓ 1623ms | http |
| 180.127.149.228:1080 | ✓ 933ms | ✓ 1202ms | ✓ 1408ms | ✓ 1936ms | ✓ 868ms | http |
| 38.180.2.107:3128 | ✓ 953ms | ✓ 1956ms | ✓ 1648ms | 否 | ✓ 1929ms | http |
| 116.80.96.110:3172 | 否 | 否 | ✓ 1529ms | ✓ 1790ms | ✓ 1646ms | http |
| 116.80.96.111:3172 | 否 | 否 | ✓ 1523ms | ✓ 1796ms | ✓ 1621ms | http |
| 101.32.244.83:8080 | ✓ 1463ms | 否 | ✓ 902ms | ✓ 1205ms | ✓ 1167ms | http |
| 121.43.196.213:8222 | ✓ 942ms | ✓ 1019ms | ✓ 844ms | ✓ 1036ms | ✓ 836ms | http |
| 121.43.196.210:8222 | ✓ 918ms | ✓ 1021ms | ✓ 866ms | ✓ 1080ms | ✓ 825ms | http |
| 114.55.226.123:10086 | ✓ 1023ms | ✓ 1509ms | ✓ 972ms | ✓ 1258ms | ✓ 1006ms | http |
| 106.117.208.101:7890 | ✓ 970ms | 否 | ✓ 1050ms | ✓ 1226ms | ✓ 988ms | http |
| 45.140.147.155:1082 | ✓ 630ms | 否 | ✓ 1069ms | 否 | ✓ 1494ms | http |
| 121.230.9.248:1080 | ✓ 1441ms | 否 | ✓ 1786ms | 否 | ✓ 1889ms | http |
| 152.42.213.210:443 | ✓ 846ms | 否 | ✓ 942ms | 否 | ✓ 1806ms | http |
| 172.212.68.37:3128 | ✓ 392ms | ✓ 1711ms | ✓ 1658ms | ✓ 1754ms | ✓ 1190ms | http |
| 121.42.162.62:1110 | ✓ 1356ms | ✓ 1310ms | ✓ 1258ms | ✓ 1532ms | ✓ 1363ms | http |
| 45.136.130.239:8443 | 否 | ✓ 1205ms | 否 | ✓ 1820ms | ✓ 679ms | http |
| 116.80.96.108:3172 | ✓ 1501ms | 否 | ✓ 1470ms | ✓ 1796ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1781ms | 否 | ✓ 1735ms | ✓ 1473ms | ✓ 1271ms | http |
| 150.107.140.238:3128 | ✓ 1816ms | 否 | ✓ 1643ms | ✓ 1433ms | ✓ 1360ms | http |
| 221.122.91.36:11195 | 否 | ✓ 1504ms | ✓ 1070ms | ✓ 1565ms | ✓ 1070ms | http |
| 121.230.9.160:1080 | 否 | 否 | ✓ 1268ms | ✓ 1790ms | ✓ 991ms | http |
| 42.84.157.30:10808 | ✓ 944ms | ✓ 1212ms | ✓ 1040ms | ✓ 1226ms | ✓ 964ms | http |
| 194.163.183.242:3128 | ✓ 1028ms | 否 | ✓ 903ms | ✓ 1790ms | ✓ 1376ms | http |
| 152.70.84.108:8080 | ✓ 747ms | 否 | ✓ 949ms | ✓ 985ms | ✓ 704ms | http |
| 159.223.42.219:3128 | ✓ 1898ms | 否 | ✓ 803ms | ✓ 1029ms | ✓ 857ms | http |
| 88.80.150.82:8080 | ✓ 1211ms | 否 | 否 | ✓ 1985ms | ✓ 1618ms | https |
| 160.250.5.22:1 | ✓ 1451ms | 否 | ✓ 1662ms | ✓ 1648ms | ✓ 1083ms | http |

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
