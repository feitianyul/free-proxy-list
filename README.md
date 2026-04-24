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

最后更新：2026-04-24 00:41:37 UTC（2026-04-24 08:41:37 UTC+8）

**代理总数：62**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1658ms | ✓ 1313ms | ✓ 899ms | ✓ 1145ms | ✓ 895ms | http |
| 46.101.95.183:8888 | ✓ 1123ms | ✓ 1624ms | ✓ 895ms | ✓ 1557ms | ✓ 1934ms | http |
| 113.160.132.26:8080 | ✓ 1795ms | ✓ 1627ms | ✓ 1545ms | ✓ 1419ms | ✓ 1127ms | http |
| 130.61.174.200:1080 | ✓ 1605ms | ✓ 1568ms | 否 | ✓ 1266ms | ✓ 986ms | http |
| 212.58.132.5:8888 | ✓ 1151ms | 否 | ✓ 1722ms | ✓ 1921ms | ✓ 1293ms | http |
| 168.110.52.228:3128 | ✓ 795ms | ✓ 1272ms | ✓ 1380ms | ✓ 981ms | ✓ 789ms | http |
| 218.108.131.186:17890 | ✓ 1168ms | 否 | ✓ 1003ms | 否 | ✓ 1036ms | http |
| 115.231.181.40:8128 | ✓ 1034ms | ✓ 1453ms | ✓ 1045ms | ✓ 1485ms | ✓ 1210ms | http |
| 45.153.231.229:8080 | ✓ 833ms | ✓ 1812ms | ✓ 1712ms | ✓ 1690ms | 否 | http |
| 120.92.108.86:7890 | ✓ 1614ms | 否 | ✓ 1868ms | ✓ 1814ms | ✓ 1530ms | http |
| 103.82.23.118:5261 | ✓ 1391ms | 否 | ✓ 1888ms | 否 | ✓ 1849ms | http |
| 84.47.150.125:1080 | ✓ 1012ms | 否 | ✓ 1967ms | ✓ 1897ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1163ms | 否 | ✓ 1165ms | ✓ 1538ms | 否 | http |
| 47.85.51.197:1080 | ✓ 108ms | ✓ 815ms | ✓ 347ms | ✓ 898ms | ✓ 1026ms | http |
| 62.113.119.14:8080 | ✓ 571ms | ✓ 1492ms | ✓ 986ms | ✓ 1509ms | ✓ 1148ms | http |
| 223.84.151.86:30005 | 否 | 否 | ✓ 1800ms | ✓ 1922ms | ✓ 1772ms | http |
| 20.127.128.70:8080 | ✓ 811ms | ✓ 1782ms | ✓ 470ms | ✓ 1463ms | 否 | http |
| 91.233.223.147:3128 | ✓ 1556ms | ✓ 1845ms | ✓ 1150ms | 否 | ✓ 1945ms | http |
| 35.225.22.61:80 | ✓ 881ms | ✓ 1258ms | 否 | ✓ 1150ms | ✓ 1321ms | http |
| 177.93.132.244:3128 | ✓ 648ms | 否 | ✓ 855ms | 否 | ✓ 1634ms | http |
| 85.190.99.143:443 | ✓ 685ms | 否 | ✓ 852ms | ✓ 1988ms | ✓ 1447ms | http |
| 91.99.15.45:2095 | ✓ 711ms | ✓ 1893ms | ✓ 1123ms | 否 | 否 | http |
| 45.140.147.155:1081 | ✓ 1284ms | ✓ 1922ms | 否 | 否 | ✓ 1345ms | http |
| 92.113.149.172:1080 | ✓ 1008ms | 否 | ✓ 1153ms | ✓ 1875ms | ✓ 1850ms | http |
| 45.140.147.155:1082 | ✓ 1303ms | ✓ 1849ms | 否 | 否 | ✓ 1413ms | http |
| 120.92.212.16:8890 | ✓ 1255ms | ✓ 1393ms | ✓ 1575ms | 否 | 否 | http |
| 120.92.212.16:7890 | ✓ 1163ms | ✓ 1747ms | ✓ 1397ms | 否 | ✓ 1178ms | http |
| 168.144.75.9:3128 | ✓ 1263ms | 否 | ✓ 1469ms | ✓ 1807ms | ✓ 1052ms | http |
| 155.212.188.205:8080 | ✓ 835ms | 否 | ✓ 1849ms | 否 | ✓ 1943ms | http |
| 14.143.222.113:57788 | ✓ 1321ms | 否 | ✓ 1759ms | ✓ 1369ms | 否 | http |
| 210.223.44.230:3128 | 否 | ✓ 1205ms | ✓ 1432ms | ✓ 1115ms | 否 | http |
| 146.56.185.39:10900 | ✓ 1820ms | 否 | 否 | ✓ 1603ms | ✓ 1374ms | http |
| 105.159.172.5:4145 | ✓ 1761ms | ✓ 1859ms | ✓ 1572ms | ✓ 1824ms | 否 | http |
| 8.211.166.184:8081 | ✓ 664ms | 否 | ✓ 711ms | ✓ 1109ms | ✓ 850ms | http |
| 218.77.106.10:10150 | ✓ 1052ms | ✓ 1403ms | ✓ 1173ms | ✓ 1440ms | ✓ 1147ms | http |
| 91.193.240.157:9877 | ✓ 970ms | 否 | ✓ 913ms | 否 | ✓ 1597ms | http |
| 152.42.177.32:8888 | ✓ 1171ms | 否 | ✓ 1519ms | ✓ 1508ms | ✓ 1538ms | http |
| 150.107.140.238:3128 | ✓ 1855ms | 否 | ✓ 1046ms | ✓ 1652ms | ✓ 1135ms | http |
| 45.88.0.115:3128 | ✓ 493ms | ✓ 1472ms | ✓ 1407ms | 否 | ✓ 1604ms | http |
| 194.147.90.23:3128 | ✓ 976ms | 否 | ✓ 1402ms | 否 | ✓ 1942ms | http |
| 45.88.0.113:3128 | ✓ 418ms | ✓ 1212ms | ✓ 870ms | ✓ 1183ms | 否 | http |
| 45.88.0.99:3128 | ✓ 365ms | ✓ 1357ms | ✓ 354ms | ✓ 1178ms | ✓ 906ms | http |
| 47.105.98.23:3128 | 否 | ✓ 1362ms | ✓ 1158ms | ✓ 1445ms | ✓ 1142ms | http |
| 38.242.223.121:3128 | 否 | ✓ 1282ms | ✓ 1813ms | ✓ 1972ms | ✓ 1641ms | http |
| 34.101.184.164:3128 | ✓ 1708ms | 否 | ✓ 1285ms | ✓ 1545ms | ✓ 1559ms | http |
| 15.204.151.141:3128 | ✓ 1145ms | 否 | ✓ 1009ms | 否 | ✓ 1385ms | http |
| 8.219.195.129:1080 | ✓ 1160ms | ✓ 1915ms | ✓ 1133ms | ✓ 1271ms | ✓ 1014ms | http |
| 137.59.47.73:3128 | ✓ 1130ms | 否 | 否 | ✓ 1521ms | ✓ 1459ms | http |
| 45.140.147.82:1081 | ✓ 536ms | 否 | ✓ 1571ms | ✓ 1966ms | 否 | http |
| 45.88.0.111:3128 | ✓ 424ms | 否 | ✓ 362ms | ✓ 1920ms | ✓ 1452ms | http |
| 213.220.62.63:3128 | ✓ 433ms | ✓ 1185ms | ✓ 1058ms | 否 | ✓ 1494ms | http |
| 60.249.94.208:3128 | ✓ 1196ms | ✓ 1431ms | ✓ 1014ms | ✓ 1461ms | ✓ 1214ms | http |
| 138.124.99.216:8888 | ✓ 879ms | ✓ 1637ms | ✓ 1765ms | ✓ 1471ms | ✓ 1140ms | http |
| 82.148.18.242:443 | ✓ 656ms | ✓ 1617ms | ✓ 1726ms | 否 | 否 | http |
| 152.70.91.193:40000 | ✓ 1554ms | 否 | ✓ 1493ms | 否 | ✓ 1755ms | http |
| 149.248.76.55:10000 | ✓ 631ms | ✓ 1511ms | ✓ 892ms | ✓ 1666ms | ✓ 1211ms | http |
| 146.19.56.212:40002 | ✓ 776ms | 否 | ✓ 372ms | ✓ 1681ms | ✓ 1346ms | http |
| 47.84.73.61:1080 | 否 | ✓ 1896ms | ✓ 904ms | ✓ 1291ms | ✓ 1043ms | http |
| 13.41.196.179:9002 | ✓ 1792ms | 否 | ✓ 1286ms | 否 | ✓ 1710ms | http |
| 103.187.146.151:3128 | ✓ 1462ms | 否 | ✓ 1079ms | ✓ 1316ms | ✓ 1046ms | http |
| 208.87.243.199:7878 | ✓ 1495ms | ✓ 1968ms | 否 | 否 | ✓ 1122ms | http |
| 45.88.0.117:3128 | ✓ 883ms | ✓ 1201ms | ✓ 1079ms | ✓ 1885ms | ✓ 1513ms | http |

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
