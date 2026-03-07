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

最后更新：2026-03-07 09:29:45 UTC（2026-03-07 17:29:45 UTC+8）

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
| 205.209.118.30:3138 | ✓ 790ms | ✓ 1007ms | ✓ 627ms | ✓ 940ms | ✓ 695ms | http |
| 136.49.39.94:8888 | ✓ 886ms | 否 | 否 | ✓ 1573ms | ✓ 1062ms | http |
| 1.231.81.166:3128 | ✓ 1565ms | 否 | ✓ 1088ms | ✓ 1097ms | ✓ 879ms | http |
| 61.72.221.94:3128 | ✓ 1561ms | ✓ 1712ms | ✓ 1869ms | ✓ 1970ms | 否 | http |
| 1.225.116.115:1080 | ✓ 1190ms | ✓ 1647ms | ✓ 1142ms | ✓ 1573ms | ✓ 1241ms | http |
| 125.128.12.144:3128 | ✓ 1056ms | ✓ 1607ms | 否 | ✓ 1888ms | ✓ 1026ms | http |
| 217.76.245.80:999 | ✓ 869ms | ✓ 1492ms | ✓ 1104ms | ✓ 1482ms | ✓ 1118ms | http |
| 14.56.107.244:3128 | 否 | 否 | ✓ 1180ms | ✓ 1786ms | ✓ 1575ms | http |
| 46.183.25.8:443 | ✓ 1399ms | 否 | ✓ 1699ms | ✓ 1974ms | 否 | http |
| 103.113.70.189:1081 | 否 | ✓ 896ms | 否 | ✓ 1321ms | ✓ 1680ms | http |
| 159.223.42.219:3128 | 否 | 否 | ✓ 895ms | ✓ 1285ms | ✓ 1013ms | http |
| 167.172.69.123:8080 | 否 | 否 | ✓ 1157ms | ✓ 1626ms | ✓ 998ms | http |
| 61.72.221.194:3128 | 否 | 否 | ✓ 1542ms | ✓ 1293ms | ✓ 1305ms | http |
| 178.236.245.59:3128 | ✓ 866ms | 否 | ✓ 1491ms | 否 | ✓ 1772ms | http |
| 178.236.245.17:3128 | ✓ 746ms | 否 | ✓ 1611ms | 否 | ✓ 1870ms | http |
| 116.80.82.224:3172 | ✓ 1659ms | 否 | ✓ 1717ms | 否 | ✓ 1844ms | http |
| 193.168.173.136:443 | ✓ 722ms | 否 | ✓ 1441ms | ✓ 1624ms | 否 | http |
| 46.249.103.192:443 | ✓ 749ms | 否 | ✓ 1687ms | ✓ 1808ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1250ms | ✓ 1436ms | ✓ 1150ms | 否 | ✓ 1154ms | http |
| 138.124.53.25:7443 | ✓ 525ms | 否 | 否 | ✓ 1434ms | ✓ 1290ms | http |
| 14.225.217.30:7890 | ✓ 1543ms | 否 | ✓ 968ms | ✓ 1575ms | 否 | http |
| 167.172.69.123:80 | ✓ 1589ms | 否 | ✓ 1280ms | ✓ 1241ms | ✓ 989ms | http |
| 120.92.212.16:8890 | ✓ 1162ms | 否 | ✓ 1126ms | 否 | ✓ 1166ms | http |
| 150.107.140.238:3128 | ✓ 1732ms | 否 | ✓ 1122ms | ✓ 1434ms | ✓ 1071ms | http |
| 101.43.255.96:80 | 否 | ✓ 1817ms | 否 | ✓ 1540ms | ✓ 1509ms | http |
| 185.115.74.185:8080 | ✓ 730ms | ✓ 1694ms | ✓ 1288ms | 否 | 否 | http |
| 91.193.240.157:9877 | ✓ 852ms | 否 | ✓ 1161ms | 否 | ✓ 1692ms | http |
| 81.70.169.194:80 | ✓ 1181ms | 否 | ✓ 1554ms | ✓ 1588ms | 否 | http |
| 190.9.109.205:999 | ✓ 663ms | ✓ 1418ms | ✓ 1058ms | ✓ 1309ms | ✓ 1187ms | http |
| 190.9.109.199:999 | ✓ 682ms | ✓ 1542ms | ✓ 1221ms | ✓ 1177ms | 否 | http |
| 89.185.85.138:1080 | ✓ 420ms | 否 | ✓ 925ms | ✓ 1727ms | ✓ 1245ms | http |
| 35.225.22.61:80 | ✓ 327ms | 否 | ✓ 949ms | ✓ 1013ms | ✓ 910ms | http |
| 159.89.31.62:8080 | ✓ 370ms | 否 | ✓ 427ms | ✓ 1370ms | ✓ 962ms | http |
| 88.80.150.82:8080 | ✓ 1197ms | 否 | ✓ 1394ms | 否 | ✓ 1884ms | https |
| 103.215.36.88:17890 | ✓ 1196ms | ✓ 1544ms | 否 | ✓ 1495ms | ✓ 1133ms | http |
| 103.215.36.88:16938 | 否 | 否 | ✓ 1801ms | ✓ 1912ms | ✓ 1526ms | http |
| 61.72.221.234:3128 | ✓ 1599ms | ✓ 1606ms | ✓ 805ms | 否 | 否 | http |
| 61.72.110.94:3128 | 否 | ✓ 1763ms | ✓ 1105ms | 否 | ✓ 983ms | http |
| 125.128.12.14:3128 | ✓ 1616ms | 否 | ✓ 1139ms | 否 | ✓ 1053ms | http |
| 121.128.121.54:3128 | ✓ 1998ms | 否 | ✓ 756ms | ✓ 1343ms | 否 | http |
| 61.72.110.54:3128 | ✓ 1602ms | 否 | ✓ 1014ms | ✓ 1576ms | ✓ 1012ms | http |
| 85.9.195.140:1080 | ✓ 411ms | 否 | ✓ 298ms | ✓ 1690ms | ✓ 1433ms | http |
| 91.107.175.112:10801 | ✓ 1996ms | 否 | 否 | ✓ 1939ms | ✓ 1564ms | http |
| 188.132.141.249:443 | ✓ 1125ms | 否 | ✓ 1573ms | 否 | ✓ 1777ms | http |
| 103.84.95.54:7890 | ✓ 1330ms | 否 | ✓ 1346ms | ✓ 1358ms | 否 | http |
| 45.140.147.82:1082 | ✓ 536ms | ✓ 1468ms | ✓ 402ms | ✓ 1302ms | 否 | http |
| 14.225.222.247:7890 | 否 | 否 | ✓ 1998ms | ✓ 1365ms | ✓ 1754ms | http |
| 34.101.184.164:3128 | ✓ 1150ms | 否 | 否 | ✓ 1799ms | ✓ 1388ms | http |
| 14.225.211.139:7890 | ✓ 1438ms | 否 | 否 | ✓ 1407ms | ✓ 1233ms | http |
| 192.166.82.55:1080 | ✓ 1811ms | ✓ 1880ms | ✓ 1694ms | ✓ 1779ms | ✓ 1323ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1374ms | ✓ 1497ms | ✓ 1475ms | ✓ 1126ms | http |
| 150.136.153.231:80 | ✓ 272ms | ✓ 1540ms | 否 | 否 | ✓ 1416ms | http |
| 202.58.77.77:1111 | 否 | 否 | ✓ 1682ms | ✓ 1783ms | ✓ 1898ms | http |
| 101.32.244.83:8080 | 否 | 否 | ✓ 1202ms | ✓ 1657ms | ✓ 1525ms | http |
| 121.43.196.210:8222 | ✓ 1082ms | ✓ 1242ms | ✓ 1074ms | ✓ 1369ms | ✓ 1056ms | http |
| 121.43.196.213:8222 | ✓ 1071ms | ✓ 1358ms | ✓ 994ms | ✓ 1362ms | ✓ 1062ms | http |
| 114.55.226.123:10086 | ✓ 1250ms | 否 | ✓ 1221ms | ✓ 1503ms | ✓ 1281ms | http |
| 194.59.204.87:9080 | ✓ 1046ms | ✓ 1495ms | ✓ 1061ms | 否 | 否 | http |
| 51.250.37.15:6666 | ✓ 1286ms | ✓ 1870ms | ✓ 1398ms | 否 | 否 | http |
| 4.213.222.169:3128 | ✓ 1985ms | ✓ 1822ms | ✓ 1457ms | ✓ 1607ms | ✓ 1502ms | http |
| 103.139.138.194:3128 | 否 | 否 | ✓ 1993ms | ✓ 1898ms | ✓ 1377ms | http |
| 162.248.165.72:1080 | ✓ 796ms | 否 | ✓ 890ms | ✓ 1453ms | 否 | http |
| 168.235.110.63:3128 | ✓ 1281ms | ✓ 1924ms | ✓ 1018ms | ✓ 973ms | ✓ 703ms | http |
| 45.129.141.143:3128 | ✓ 890ms | 否 | ✓ 1610ms | ✓ 1832ms | 否 | http |
| 211.171.114.154:3128 | 否 | ✓ 1664ms | ✓ 1585ms | 否 | ✓ 1365ms | http |
| 103.82.23.118:5185 | ✓ 1467ms | 否 | ✓ 1398ms | 否 | ✓ 1569ms | http |
| 14.225.222.164:7890 | ✓ 1667ms | ✓ 1893ms | 否 | 否 | ✓ 1715ms | http |
| 107.172.125.217:3128 | ✓ 1363ms | 否 | ✓ 1894ms | ✓ 1235ms | ✓ 789ms | http |
| 172.212.68.37:3128 | ✓ 278ms | 否 | ✓ 741ms | ✓ 1431ms | ✓ 890ms | http |
| 45.136.198.40:3128 | ✓ 696ms | ✓ 1482ms | 否 | 否 | ✓ 1667ms | http |
| 103.215.36.88:16431 | ✓ 1202ms | 否 | ✓ 1942ms | 否 | ✓ 1615ms | http |
| 103.215.36.88:18378 | ✓ 1622ms | 否 | ✓ 1372ms | ✓ 1990ms | ✓ 1254ms | http |
| 210.223.44.230:3128 | ✓ 1774ms | ✓ 1527ms | 否 | ✓ 1174ms | ✓ 909ms | http |
| 51.158.61.240:13128 | 否 | ✓ 1871ms | ✓ 1815ms | ✓ 1878ms | 否 | http |
| 207.254.71.62:8088 | ✓ 557ms | 否 | 否 | ✓ 1560ms | ✓ 1860ms | http |
| 144.31.25.69:21064 | ✓ 982ms | 否 | ✓ 849ms | 否 | ✓ 1994ms | http |
| 106.14.203.63:3333 | ✓ 1064ms | ✓ 1271ms | 否 | 否 | ✓ 1040ms | http |

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
