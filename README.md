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

最后更新：2026-04-18 14:01:24 UTC（2026-04-18 22:01:24 UTC+8）

**代理总数：83**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | 否 | ✓ 1265ms | ✓ 989ms | ✓ 1157ms | ✓ 974ms | http |
| 194.87.130.16:1080 | ✓ 644ms | 否 | 否 | ✓ 1950ms | ✓ 1288ms | http |
| 81.30.156.115:8080 | ✓ 450ms | ✓ 1970ms | ✓ 1396ms | ✓ 1947ms | ✓ 1767ms | http |
| 45.67.202.178:1080 | ✓ 1093ms | 否 | ✓ 1769ms | 否 | ✓ 1948ms | http |
| 113.160.132.26:8080 | ✓ 1957ms | 否 | 否 | ✓ 1528ms | ✓ 1501ms | http |
| 149.51.42.10:8080 | ✓ 490ms | ✓ 1314ms | 否 | ✓ 1249ms | 否 | http |
| 185.138.116.150:8080 | ✓ 848ms | ✓ 1568ms | ✓ 621ms | ✓ 1428ms | ✓ 1069ms | http |
| 188.246.224.49:7890 | ✓ 753ms | ✓ 1255ms | ✓ 633ms | 否 | ✓ 1276ms | http |
| 45.140.147.82:1081 | ✓ 660ms | 否 | ✓ 469ms | ✓ 1367ms | ✓ 1077ms | http |
| 46.101.95.183:8888 | ✓ 821ms | 否 | ✓ 563ms | ✓ 1684ms | 否 | http |
| 152.42.208.139:8118 | ✓ 924ms | 否 | ✓ 1279ms | ✓ 1271ms | ✓ 1006ms | http |
| 162.19.253.202:8443 | ✓ 1845ms | 否 | ✓ 923ms | 否 | ✓ 1769ms | http |
| 14.247.76.52:8080 | ✓ 1169ms | 否 | ✓ 1204ms | ✓ 1493ms | ✓ 1213ms | http |
| 117.236.124.166:3128 | ✓ 1405ms | 否 | ✓ 1751ms | 否 | ✓ 1849ms | http |
| 115.231.181.40:8128 | ✓ 1208ms | ✓ 1358ms | ✓ 1139ms | 否 | 否 | http |
| 35.225.22.61:80 | ✓ 507ms | 否 | ✓ 1036ms | ✓ 1117ms | ✓ 841ms | http |
| 45.76.207.177:40000 | ✓ 1310ms | 否 | ✓ 858ms | ✓ 1299ms | ✓ 1212ms | http |
| 84.47.150.125:1080 | 否 | 否 | ✓ 1722ms | ✓ 1967ms | ✓ 1570ms | http |
| 91.99.15.45:2095 | ✓ 678ms | ✓ 1561ms | ✓ 1726ms | ✓ 1721ms | ✓ 1672ms | http |
| 45.12.151.226:2829 | ✓ 1546ms | ✓ 1971ms | 否 | 否 | ✓ 1961ms | http |
| 20.210.39.153:8561 | ✓ 734ms | ✓ 1453ms | ✓ 712ms | ✓ 1915ms | 否 | http |
| 149.51.42.10:3128 | ✓ 820ms | ✓ 1192ms | 否 | ✓ 1671ms | 否 | http |
| 162.240.154.26:3128 | ✓ 1005ms | 否 | ✓ 1751ms | 否 | ✓ 1066ms | http |
| 150.249.255.91:3128 | ✓ 1325ms | 否 | ✓ 669ms | 否 | ✓ 1877ms | http |
| 20.27.11.248:8561 | ✓ 1312ms | ✓ 1475ms | ✓ 851ms | ✓ 989ms | ✓ 917ms | http |
| 20.210.76.175:8561 | ✓ 1314ms | ✓ 1298ms | ✓ 1046ms | ✓ 1219ms | ✓ 824ms | http |
| 47.93.216.160:1081 | ✓ 1057ms | ✓ 1411ms | ✓ 1134ms | ✓ 1411ms | ✓ 1136ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1860ms | ✓ 1629ms | ✓ 1227ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1829ms | ✓ 1643ms | ✓ 1930ms | ✓ 1333ms | http |
| 116.58.161.203:26021 | ✓ 1404ms | 否 | ✓ 1160ms | ✓ 1323ms | 否 | http |
| 20.78.26.206:8561 | ✓ 1648ms | ✓ 1207ms | ✓ 1006ms | ✓ 1132ms | ✓ 988ms | http |
| 194.104.9.38:3128 | ✓ 390ms | 否 | ✓ 685ms | ✓ 1803ms | ✓ 1855ms | http |
| 45.140.147.82:1082 | ✓ 404ms | 否 | ✓ 670ms | ✓ 1723ms | 否 | http |
| 213.32.85.26:3128 | ✓ 631ms | 否 | ✓ 653ms | 否 | ✓ 1558ms | http |
| 20.27.15.49:8561 | ✓ 1663ms | 否 | ✓ 666ms | ✓ 1073ms | ✓ 979ms | http |
| 210.223.44.230:3128 | ✓ 1736ms | ✓ 1884ms | 否 | 否 | ✓ 1094ms | http |
| 161.97.184.191:8080 | ✓ 913ms | ✓ 1622ms | ✓ 1183ms | 否 | 否 | http |
| 164.163.42.25:10000 | ✓ 1110ms | 否 | ✓ 1110ms | 否 | ✓ 1761ms | http |
| 212.58.132.5:8888 | ✓ 1526ms | 否 | ✓ 1677ms | ✓ 1503ms | ✓ 1370ms | http |
| 108.181.201.118:1234 | ✓ 826ms | ✓ 1144ms | 否 | 否 | ✓ 1100ms | http |
| 195.26.224.49:3128 | ✓ 839ms | 否 | ✓ 1015ms | ✓ 1853ms | 否 | http |
| 20.18.193.135:8561 | ✓ 1868ms | 否 | 否 | ✓ 1975ms | ✓ 1489ms | http |
| 62.113.119.14:8080 | ✓ 1282ms | 否 | ✓ 777ms | 否 | ✓ 1395ms | http |
| 47.76.189.189:8899 | ✓ 1317ms | 否 | ✓ 1279ms | ✓ 1530ms | ✓ 1308ms | http |
| 101.32.244.83:8080 | ✓ 1554ms | 否 | ✓ 1432ms | ✓ 1770ms | ✓ 1560ms | http |
| 121.43.196.213:8222 | 否 | ✓ 1281ms | ✓ 1055ms | ✓ 1317ms | ✓ 1088ms | http |
| 121.43.196.210:8222 | 否 | ✓ 1446ms | ✓ 1096ms | ✓ 1336ms | ✓ 1167ms | http |
| 114.55.226.123:10086 | ✓ 1271ms | 否 | ✓ 1316ms | ✓ 1575ms | ✓ 1305ms | http |
| 218.108.131.186:17890 | 否 | ✓ 1258ms | ✓ 1071ms | ✓ 1330ms | ✓ 1044ms | http |
| 20.27.15.111:8561 | 否 | ✓ 1745ms | ✓ 1162ms | ✓ 1472ms | ✓ 1049ms | http |
| 20.78.118.91:8561 | ✓ 768ms | ✓ 1633ms | ✓ 985ms | ✓ 1036ms | ✓ 815ms | http |
| 20.210.76.104:8561 | ✓ 1301ms | ✓ 1822ms | ✓ 1430ms | ✓ 1823ms | ✓ 1391ms | http |
| 20.210.76.178:8561 | 否 | 否 | ✓ 1332ms | ✓ 1617ms | ✓ 1378ms | http |
| 45.140.147.155:1082 | ✓ 530ms | 否 | ✓ 1428ms | 否 | ✓ 1343ms | http |
| 101.32.243.189:80 | ✓ 1399ms | 否 | ✓ 1498ms | 否 | ✓ 1734ms | http |
| 177.93.132.244:3128 | ✓ 666ms | 否 | ✓ 1766ms | 否 | ✓ 1616ms | http |
| 133.18.123.225:26021 | 否 | 否 | ✓ 749ms | ✓ 1114ms | ✓ 1000ms | http |
| 94.131.118.39:1082 | ✓ 1674ms | 否 | ✓ 1200ms | 否 | ✓ 1705ms | http |
| 120.92.212.16:7890 | ✓ 1908ms | ✓ 1748ms | ✓ 1164ms | ✓ 1705ms | 否 | http |
| 89.111.174.221:8080 | 否 | ✓ 1872ms | ✓ 1305ms | 否 | ✓ 1634ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1251ms | ✓ 1393ms | ✓ 1237ms | http |
| 223.84.151.86:30005 | 否 | ✓ 1516ms | ✓ 1377ms | ✓ 1727ms | ✓ 1541ms | http |
| 20.27.13.35:8561 | ✓ 1837ms | ✓ 1268ms | ✓ 697ms | ✓ 1080ms | ✓ 1024ms | http |
| 217.76.245.80:999 | ✓ 1779ms | 否 | ✓ 1568ms | ✓ 1445ms | ✓ 1249ms | http |
| 20.27.14.220:8561 | ✓ 658ms | 否 | ✓ 731ms | ✓ 1077ms | ✓ 795ms | http |
| 194.67.127.23:10808 | 否 | 否 | ✓ 1431ms | ✓ 1736ms | ✓ 1503ms | http |
| 43.132.188.134:443 | ✓ 1459ms | ✓ 1454ms | 否 | ✓ 1295ms | ✓ 1834ms | http |
| 144.31.27.49:1080 | ✓ 594ms | ✓ 1627ms | 否 | 否 | ✓ 1653ms | http |
| 103.85.113.66:9999 | ✓ 1111ms | 否 | ✓ 861ms | ✓ 1804ms | 否 | http |
| 147.45.214.210:1080 | 否 | ✓ 1834ms | 否 | ✓ 1737ms | ✓ 1872ms | http |
| 94.158.219.111:3128 | ✓ 700ms | ✓ 1628ms | ✓ 629ms | 否 | ✓ 1654ms | http |
| 20.127.128.70:8080 | ✓ 537ms | 否 | 否 | ✓ 1123ms | ✓ 853ms | http |
| 103.82.23.118:5205 | ✓ 1735ms | 否 | ✓ 1641ms | 否 | ✓ 1816ms | http |
| 78.11.96.22:8888 | ✓ 938ms | 否 | ✓ 940ms | ✓ 1399ms | ✓ 1253ms | http |
| 138.124.99.216:8888 | 否 | ✓ 1775ms | ✓ 1521ms | ✓ 1833ms | ✓ 1481ms | http |
| 84.47.150.126:1080 | 否 | 否 | ✓ 1577ms | ✓ 1836ms | ✓ 1208ms | http |
| 185.114.73.2:1080 | ✓ 1911ms | 否 | ✓ 1622ms | 否 | ✓ 1490ms | http |
| 82.114.228.67:1080 | ✓ 1137ms | ✓ 1426ms | ✓ 1634ms | 否 | 否 | http |
| 116.80.48.16:7777 | ✓ 1721ms | 否 | ✓ 1713ms | 否 | ✓ 1861ms | http |
| 61.52.131.172:8443 | ✓ 1095ms | 否 | ✓ 1110ms | ✓ 1448ms | ✓ 1569ms | http |
| 8.219.195.129:1080 | ✓ 1559ms | ✓ 1906ms | ✓ 1107ms | ✓ 1296ms | ✓ 1056ms | http |
| 160.250.5.22:1 | 否 | 否 | ✓ 1787ms | ✓ 1607ms | ✓ 1366ms | http |
| 211.95.152.50:45046 | 否 | ✓ 1735ms | ✓ 1498ms | ✓ 1538ms | ✓ 1363ms | http |

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
