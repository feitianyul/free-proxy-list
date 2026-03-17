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

最后更新：2026-03-17 12:36:17 UTC（2026-03-17 20:36:17 UTC+8）

**代理总数：52**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 52 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 52 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 202.155.12.161:443 | ✓ 1369ms | 否 | ✓ 1326ms | ✓ 1138ms | 否 | http |
| 113.160.132.26:8080 | 否 | ✓ 1736ms | ✓ 971ms | ✓ 1273ms | 否 | http |
| 147.161.210.140:8800 | ✓ 1338ms | 否 | ✓ 920ms | ✓ 1389ms | ✓ 1070ms | http |
| 91.107.148.58:53967 | ✓ 1083ms | 否 | ✓ 1487ms | 否 | ✓ 1912ms | http |
| 101.47.73.135:3128 | ✓ 1152ms | 否 | ✓ 1134ms | 否 | ✓ 1167ms | http |
| 194.5.212.40:8080 | ✓ 1238ms | 否 | ✓ 1316ms | ✓ 1826ms | ✓ 1425ms | http |
| 120.92.212.16:8890 | ✓ 994ms | ✓ 1243ms | ✓ 1072ms | 否 | ✓ 992ms | http |
| 116.80.65.81:3172 | ✓ 1696ms | 否 | 否 | ✓ 1862ms | ✓ 1687ms | http |
| 45.167.124.52:8080 | 否 | 否 | ✓ 1429ms | ✓ 1697ms | ✓ 1435ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1263ms | 否 | ✓ 1252ms | ✓ 1025ms | http |
| 62.60.177.204:34094 | 否 | ✓ 1526ms | 否 | ✓ 1272ms | ✓ 807ms | http |
| 101.43.127.100:8877 | ✓ 1826ms | ✓ 1057ms | 否 | ✓ 1204ms | 否 | http |
| 137.220.150.170:6005 | 否 | 否 | ✓ 959ms | ✓ 1170ms | ✓ 1079ms | http |
| 116.80.96.105:3172 | ✓ 1515ms | 否 | ✓ 1518ms | ✓ 1841ms | 否 | http |
| 52.74.26.202:8080 | ✓ 785ms | ✓ 1226ms | ✓ 1156ms | ✓ 1197ms | ✓ 1222ms | http |
| 103.82.23.118:5171 | 否 | 否 | ✓ 1703ms | ✓ 1625ms | ✓ 1626ms | http |
| 219.117.204.211:7799 | ✓ 1386ms | ✓ 1071ms | ✓ 542ms | ✓ 1298ms | 否 | http |
| 47.77.193.180:1080 | ✓ 1262ms | ✓ 1135ms | ✓ 508ms | ✓ 762ms | ✓ 613ms | http |
| 116.80.65.75:3172 | 否 | 否 | ✓ 1666ms | ✓ 1888ms | ✓ 1689ms | http |
| 137.220.151.110:6005 | ✓ 777ms | 否 | ✓ 1135ms | ✓ 1328ms | ✓ 858ms | http |
| 212.192.12.90:6005 | ✓ 1400ms | 否 | 否 | ✓ 1605ms | ✓ 1070ms | http |
| 116.80.49.156:3172 | ✓ 1875ms | 否 | 否 | ✓ 1961ms | ✓ 1926ms | http |
| 59.46.216.131:30001 | ✓ 1753ms | 否 | 否 | ✓ 1822ms | ✓ 1409ms | http |
| 103.113.70.189:1081 | ✓ 1571ms | 否 | 否 | ✓ 1124ms | ✓ 897ms | http |
| 150.249.255.91:3128 | 否 | 否 | ✓ 641ms | ✓ 1329ms | ✓ 1453ms | http |
| 86.53.183.16:1080 | ✓ 1274ms | 否 | ✓ 748ms | 否 | ✓ 1823ms | http |
| 137.220.150.152:6005 | ✓ 1821ms | 否 | 否 | ✓ 1120ms | ✓ 887ms | http |
| 133.242.138.34:8100 | ✓ 1995ms | ✓ 1268ms | 否 | ✓ 1980ms | ✓ 1471ms | http |
| 164.90.155.209:3128 | 否 | 否 | ✓ 792ms | ✓ 796ms | ✓ 576ms | http |
| 168.235.110.63:3128 | ✓ 864ms | 否 | ✓ 1783ms | ✓ 1116ms | 否 | http |
| 45.136.198.40:3128 | ✓ 822ms | 否 | ✓ 1049ms | 否 | ✓ 1813ms | http |
| 35.225.22.61:80 | ✓ 1412ms | 否 | ✓ 403ms | ✓ 1148ms | 否 | http |
| 120.55.163.237:10086 | ✓ 959ms | ✓ 1095ms | ✓ 988ms | ✓ 1145ms | ✓ 899ms | http |
| 212.192.13.76:6005 | ✓ 1156ms | 否 | ✓ 1339ms | ✓ 1362ms | ✓ 1023ms | http |
| 106.117.208.101:7890 | ✓ 1074ms | ✓ 1344ms | 否 | ✓ 1435ms | ✓ 1077ms | http |
| 103.124.137.189:3128 | 否 | 否 | ✓ 1674ms | ✓ 1853ms | ✓ 1426ms | http |
| 1.231.81.166:3128 | ✓ 1089ms | ✓ 1004ms | ✓ 1488ms | ✓ 1325ms | ✓ 1128ms | http |
| 165.227.5.10:8888 | ✓ 1766ms | ✓ 1630ms | 否 | ✓ 1171ms | 否 | http |
| 186.148.180.46:999 | ✓ 1167ms | ✓ 1933ms | ✓ 1497ms | 否 | ✓ 1697ms | http |
| 103.39.51.190:8080 | ✓ 1834ms | 否 | ✓ 1777ms | ✓ 1561ms | 否 | http |
| 45.168.238.193:8443 | 否 | ✓ 1775ms | ✓ 477ms | ✓ 1153ms | ✓ 874ms | http |
| 121.40.231.103:7890 | ✓ 1705ms | ✓ 1137ms | ✓ 888ms | ✓ 1143ms | ✓ 878ms | http |
| 160.250.4.245:1 | ✓ 1471ms | 否 | ✓ 1349ms | ✓ 1299ms | 否 | http |
| 106.14.203.63:3333 | 否 | ✓ 1114ms | 否 | ✓ 1416ms | ✓ 1441ms | http |
| 116.80.49.159:3172 | ✓ 1921ms | 否 | 否 | ✓ 1894ms | ✓ 1647ms | http |
| 113.176.92.71:3128 | 否 | 否 | ✓ 1626ms | ✓ 1508ms | ✓ 978ms | http |
| 85.198.96.242:3128 | ✓ 1038ms | 否 | ✓ 1800ms | ✓ 1871ms | 否 | http |
| 149.50.116.240:1080 | ✓ 1450ms | 否 | ✓ 1911ms | 否 | ✓ 1979ms | http |
| 137.220.150.104:6005 | ✓ 887ms | 否 | ✓ 804ms | ✓ 1852ms | 否 | http |
| 103.69.84.106:3131 | ✓ 1722ms | 否 | ✓ 1437ms | ✓ 1166ms | ✓ 927ms | http |
| 8.219.97.248:80 | ✓ 1138ms | 否 | ✓ 944ms | ✓ 1384ms | 否 | http |
| 61.52.131.172:8443 | ✓ 900ms | 否 | ✓ 903ms | ✓ 1175ms | ✓ 913ms | http |

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
